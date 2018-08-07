package model

import (
    "os"
    "fmt"
    "bytes"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "html"
    "strings"
    "strconv"
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

// Gets the required post id
func GetPostId(request *http.Request) float64 {
    var id float64 = 0

    if len(request.FormValue("id")) > 0 {
        maybe, oops := strconv.ParseFloat(request.FormValue("id"), 64)
        if oops == nil {
            id = maybe
        }
    }

    return id
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
            post["id"] = strconv.FormatFloat(
                rawPost.(map[string]interface{})["id"].(float64),
                'e', -1, 64)
            post["title"] = rawPost.(map[string]interface{})["title"].(string)
            post["body"] = rawPost.(map[string]interface{})["body"].(string)
            outlet = append(outlet, post)
        }
    } else {
        panic("Something went wrong")
    }

    return outlet
}

// Get the post identified by the given id
func GetPost(id float64) map[string]string {
    outlet := map[string]string {
        "id": "0",
        "title": "Not found!",
        "body": "There is no post like the one you are looking for... :(",
    }
    posts := GetPosts(0)
    found := false
    offset := 0

    for (len(posts) > 10) && (!found) {
        for _, post := range posts {
            currentId, oops := strconv.ParseFloat(post["id"], 64)
            if (oops == nil) && (currentId == id) {
                outlet = post
			  	found = true
            }
        }
        offset += 10
        posts = GetPosts(offset)
    }

    return outlet
}

// Sends a simple email
func Contact(request *http.Request) error {
    // Message formatting
    message := request.FormValue("message")
    lineFeed := "\n"
    if strings.Contains(message, "\r") {
        lineFeed = "\r\n"
    }
    message = strings.Replace(message, lineFeed, "<br>", -1)
    payload := fmt.Sprintf("{\"value1\":\"%s\",\"value2\":\"%s\",\"value3\":\"%s\"}",
        request.FormValue("name"),
        request.FormValue("email"),
        html.EscapeString(message))
    content := bytes.NewBufferString(payload)

    // IFTTT setup
    iftttUrl := "https://maker.ifttt.com/trigger/%s/with/key/%s"
    iftttKey := os.Getenv("IFTTT_KEY")
    iftttEvent := "liberdade_notified"
    iftttWebhook := fmt.Sprintf(iftttUrl, iftttEvent, iftttKey)
    _, oops := http.Post(iftttWebhook, "application/json", content)
    // BUG Invalid messages are receiving an ok answer
    return oops
}
