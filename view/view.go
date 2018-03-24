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

func ShowIndex(writer io.Writer, posts []string) {
    args := make(map[string]string)



    LoadFileWithArgs(writer, "assets/html/index.gohtml", args)
}

func ShowContactPage(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/contato.gohtml")
}

func ShowAboutPage(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/quem.gohtml")
}

func DisplayError(writer io.Writer, oops error) {
    if oops == nil {
        fmt.Fprintf(writer, "Ok!\n")
    } else {
        fmt.Fprintf(writer, "%#v\n", oops)
    }
}
