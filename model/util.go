package model

import (
    "os"
    "os/exec"
    "fmt"
    "errors"
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

// Retrieves posts from Tumblr
func GetPosts() []map[string]string {
    outlet := make([]map[string]string, 0)

    // Downloading stuff
    secret := os.Getenv("TUMBLR_SECRET")
    url := fmt.Sprintf("http://api.tumblr.com/v2/blog/%s/posts/text?api_key=%s",
                       "liberdadeorganizacao.tumblr.com",
                       secret)
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
    }

    return outlet
}

// Sends a simple email
func SendSimpleMail(recipient, message string) error {
    slackWebhook := os.Getenv("SLACK_WEBHOOK")
    cmd := exec.Command("curl",
                        "-X", "POST",
                        "-H", "Content-type: application/json",
                        "--data", fmt.Sprintf("{\"text\": \"De: %s\r\nMensagem: %s\"}",
                                              recipient, message),
                        slackWebhook)
    buffer := bytes.NewBufferString("")
    cmd.Stderr = buffer
    output, oops := cmd.Output()
    if oops == nil {
        fmt.Printf("OUTPUT: %s\n", string(output))
    } else {
        why := buffer.String()
        fmt.Printf("OUTPUT: %s\n", why)
        oops = errors.New(why)
    }

    return oops
}
