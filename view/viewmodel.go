package view

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "strings"
)

// This class will deal with creating our HTML pages by adding the necessary
// assets (like CSS and Javascript) and facilitating the body customization.
type ViewModel struct {
    // This will describe the CSS style of the page
    Style template.CSS

    // This is the footer for the page, described at assets/html/footer.html
    Header template.HTML

    // This is the footer for the page, described at assets/html/footer.html
    Footer template.HTML

    // This mapping will relate the data produced by the model to the view.
    Body template.HTML

    // This is the page offset
    Offset template.HTML

    // This is the Javascript code that will be ran in the browser.
    Script template.JS
}

// Creates a new view model.
func NewViewModel() *ViewModel {
    cssFiles := []string {
        // "pure.css",
        "master.css",
    }
    jsFiles := []string {
        "app.js",
    }
    vm := ViewModel {
        Style: template.CSS(loadCss(cssFiles)),
        Header: template.HTML(loadHeader()),
        Footer: template.HTML(loadFooter()),
        Body: template.HTML(""),
        Offset: template.HTML(""),
        Script: template.JS(loadJs(jsFiles)),
    }
    return &vm
}

func GenerateViewModel(args map[string]string) *ViewModel {
    vm := NewViewModel()

    if value, ok := args["style"]; ok {
        vm.Style = template.CSS(loadCss(strings.Split(value, " ")))
    }

    if value, ok := args["script"]; ok {
        vm.Script = template.JS(loadJs(strings.Split(value, " ")))
    }

    if value, ok := args["body"]; ok {
        vm.Body = template.HTML(value)
    }

    if value,ok := args["offset"]; ok {
        vm.Offset = template.HTML(value)
    }

    return vm
}

// Load many files to a single string
func loadLot(src string, files []string) string {
    outlet := []byte { }
    pwd := GetPwd()

    for _, file := range files {
        contents, err := ioutil.ReadFile(pwd + src + file)
        if err != nil {
            panic(err)
        } else {
            for _, content := range(contents) {
                outlet = append(outlet, content)
            }
        }
    }

    return string(outlet)
}

// Extracts the CSS path
func loadCss(files []string) string {
    return loadLot("assets/css/", files)
}

// Loads the footer HTML
func loadFooter() string {
    pwd := GetPwd()
    footer, err := ioutil.ReadFile(pwd + "assets/html/footer.html")

    if err != nil {
        fmt.Println(err)
        footer = []byte { }
    }

    return string(footer)
}

// Loads the header HTML
func loadHeader() string {
    pwd := GetPwd()
    footer, err := ioutil.ReadFile(pwd + "assets/html/header.html")

    if err != nil {
        fmt.Println(err)
        footer = []byte { }
    }

    return string(footer)
}

// Loads the Javascript asset
func loadJs(scripts []string) string {
    return loadLot("assets/js/", scripts)
}
