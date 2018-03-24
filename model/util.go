package model

import "os"
import "os/exec"
import "fmt"
import "errors"
import "bytes"

// Gets the port for releasing the server
func GetPort() string {
    port := os.Getenv("PORT")

    if len(port) == 0 {
        port = "8000"
    }

    return ":" + port
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
