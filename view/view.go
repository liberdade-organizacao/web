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

func SayHello(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/index.gohtml")
}

func TalkAboutOffice(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/office.gohtml")
}

func TalkAboutGarage(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/garage.gohtml")
}

func TalkAboutGtd(writer io.Writer) {
    LoadFileWithArgs(writer, "assets/html/persorg.gohtml", map[string]string {
        "style": "reveal.css crisjr.css",
        "script": "reveal.js",
    })
}

func DisplayError(writer io.Writer, oops error) {
    if oops == nil {
        fmt.Fprintf(writer, "Ok!\n")
    } else {
        fmt.Fprintf(writer, "%#v\n", oops)
    }
}
