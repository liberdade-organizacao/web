package view

import (
    "io"
    "os"
    "html/template"
    "fmt"
)

// Gets the current PWD
func GetPwd() string {
    codePath := "./src/github.com/liberdade-organizacao/web/"
    port := os.Getenv("PORT")

    if len(port) != 0 {
        codePath = os.Getenv("HOME") + "/"
    }

    return codePath
}

// Standard procedure to load a HTML file that does not need any customization.
func LoadFileWithoutArgs(writer io.Writer, path string) {
    htmlPath := GetPwd() + path
    templ, err := template.ParseFiles(htmlPath)
    viewModel := NewViewModel()
    err = templ.Execute(writer, viewModel)
    if err != nil {
        fmt.Printf("%#v\n", err)
    }
}

// Procedure to customize HTML files
func LoadFileWithArgs(writer io.Writer, path string, args map[string]string) {
    htmlPath := GetPwd() + path
    templ, err := template.ParseFiles(htmlPath)
    viewModel := GenerateViewModel(args)
    err = templ.Execute(writer, viewModel)
    if err != nil {
        fmt.Printf("%#v\n", err)
    }
}

// Pack posts into a HTML string
func PostsToString(posts []map[string]string) string {
    body := ""

    limit := len(posts)
    if limit > 10 {
        limit = 10
    }
    for i := 0; i < limit; i++ {
        post := posts[i]
        maybe_hr := "\n<hr/>"
        if i == limit - 1 {
            maybe_hr = ""
        }
        title := fmt.Sprintf("<h2><a href=\"/blog/post?id=%s\">%s</a></h2>",
                             post["id"], post["title"])
        body = fmt.Sprintf("%s<div class=\"content\">%s%s\n%s\n</div>\n",
                           body, title, post["body"], maybe_hr)
    }

    return body
}
