package controller

import (
    "net/http"
    "github.com/liberdade-organizacao/web/model"
    "github.com/liberdade-organizacao/web/view"
)

/*********************
 * SERVER DEFINITION *
 *********************/

 type Server struct {
     Port string
 }

 // Create a webserver for this app.
func NewServer() Server {
    server := Server {
        Port: model.GetPort(),
    }

    // Main pages
    http.HandleFunc("/", server.ShowIndexPage)
    http.HandleFunc("/index", server.ShowIndexPage)
    http.HandleFunc("/contatar", server.ContactUs)
    http.HandleFunc("/suporte", server.ShowSupportPage)
    http.HandleFunc("/blog", server.DisplayBlog)

    return server
}

// Puts the webserver to, well, serve.
func (server *Server) Serve() {
    http.ListenAndServe(server.Port, nil)
}

// Displays the main page
func (server *Server) ShowIndexPage(w http.ResponseWriter, r *http.Request) {
    view.ShowIndex(w, "")
}

func (server *Server) DisplayBlog(w http.ResponseWriter, r *http.Request) {
    offset := model.GetOffset(r)
    posts := model.GetPosts(offset)
    view.ShowBlog(w, posts, offset)
}

// Displays the garage page
func (server *Server) ShowSupportPage(w http.ResponseWriter, r *http.Request) {
    view.ShowSupport(w)
}

// Sends a message to the people in Liberdade
func (server *Server) ContactUs(w http.ResponseWriter, r *http.Request) {
    oops := model.Contact(r)
    message := "ok"

    if oops == nil {

    }

    view.ShowIndex(w, message)
}
