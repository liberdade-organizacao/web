package model

import (
    "os"
    "fmt"
    "bytes"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

// Gets the port for releasing the server
func GetPort() string {
    port := os.Getenv("PORT")

    if len(port) == 0 {
        port = "8000"
    }

    return ":" + port
}

// Gets the offset value from a request
func GetOffset(request *http.Request) int {
    offset := 0

    if len(request.FormValue("offset")) > 0 {
        fmt.Sscanf(request.FormValue("offset"), "%d", &offset)
    }

    return offset
}

// Retrieves posts from Tumblr
func GetPosts(offset int) []map[string]string {
    outlet := make([]map[string]string, 0)

    // Downloading stuff
    secret := os.Getenv("TUMBLR_SECRET")
    query := "http://api.tumblr.com/v2/blog/%s/posts/text?api_key=%s&limit=11"
    url := fmt.Sprintf(query, "liberdadeorganizacao.tumblr.com", secret)
    if offset > 0 {
        url = fmt.Sprintf("%s&offset=%d", url, offset)
    }
    rawResponse, oops := http.Get(url)
    if oops != nil {
        return outlet
    }
    defer rawResponse.Body.Close()
    content, oops := ioutil.ReadAll(rawResponse.Body)

    // Parsing response
    var payload interface{}
    oops = json.Unmarshal(content, &payload)

    if oops != nil {
        return outlet
    }

    meta := payload.(map[string]interface{})["meta"].(map[string]interface{})
    status := meta["status"].(float64)
    if status == 200 {
        posts := payload.(map[string]interface{})["response"].(map[string]interface{})["posts"].([]interface{})
        for _, rawPost := range posts {
            post := make(map[string]string)
            post["title"] = rawPost.(map[string]interface{})["title"].(string)
            post["body"] = rawPost.(map[string]interface{})["body"].(string)
            outlet = append(outlet, post)
        }
    } else {
        panic("Something went wrong")
    }

    return outlet
}

// Sends a simple email
func Contact(request *http.Request) error {
    recipient := fmt.Sprintf("%s <%s>",
                             request.FormValue("name"),
                             request.FormValue("email"))
    message := request.FormValue("message")
    content := bytes.NewBufferString(fmt.Sprintf("{\"text\": \"%s: %s\"}",
                                                 recipient, message))
    slackWebhook := os.Getenv("SLACK_WEBHOOK")
    _, oops := http.Post(slackWebhook, "application/json", content)
    return oops
}
