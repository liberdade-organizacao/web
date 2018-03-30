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

func ShowIndex(writer io.Writer, posts []map[string]string, offset int) {
    args := make(map[string]string)
    body := "<div class=\"pure-u-1 pure-u-md-1-2\">"

    // Body building
    for _, post := range posts {
        title := fmt.Sprintf("<h3 class=\"information-head\">%s</h3>",
                             post["title"])
        body = fmt.Sprintf("%s<div class=\"l-box\">%s%s</div>\n",
                           body, title, post["body"])
    }
    body = fmt.Sprintf("%s</div>\n", body)
    args["body"] = body

    // Building offset
    pagination := `<p>`
    if offset >= 1 {
        off := offset-10
        if off < 0 {
            off = 0
        }
        pagination = fmt.Sprintf(`%s<a href="/index?offset=%d">
            <i class="fa fa-chevron-left" aria-hidden="true"></i>
        </a>`, pagination, off)
    }
    if len(posts) >= 10 {
        pagination = fmt.Sprintf(`%s<a href="/index?offset=%d">
            <i class="fa fa-chevron-right" aria-hidden="true"></i>
        </a>`, pagination, offset+10)
    }
    pagination = fmt.Sprintf("%s</p>", pagination)
    args["offset"] = pagination

    LoadFileWithArgs(writer, "assets/html/index.gohtml", args)
}

func ShowAboutPage(writer io.Writer, message string) {
    if len(message) == 0 {
        LoadFileWithoutArgs(writer, "assets/html/contato.gohtml")
    } else {
        script := "ok.js"
        if message != "ok" {
            script = "not-ok.js"
        }
        args := make(map[string]string)
        args["script"] = script
        LoadFileWithArgs(writer, "assets/html/contato.gohtml", args)
    }
}

func ShowSupportPage(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/suporte.gohtml")
}

func DisplayError(writer io.Writer, oops error) {
    if oops == nil {
        fmt.Fprintf(writer, "Ok!\n")
    } else {
        fmt.Fprintf(writer, "%#v\n", oops)
    }
}
