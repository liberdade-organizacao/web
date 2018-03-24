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
    http.HandleFunc("/", server.ShowIndex)
    http.HandleFunc("/contato", server.ShowContactPage)
    http.HandleFunc("/quem", server.ShowAboutPage)

    return server
}

// Puts the webserver to, well, serve.
func (server *Server) Serve() {
    http.ListenAndServe(server.Port, nil)
}

// Displays the main page
func (server *Server) ShowIndex(w http.ResponseWriter, r *http.Request) {
    // TODO Implement pagination logic
    posts := model.GetPosts()
    view.ShowIndex(w, posts)
}

// Displays the office page
func (server *Server) ShowContactPage(w http.ResponseWriter, r *http.Request) {
    view.ShowContactPage(w)
}

// Displays the garage page
func (server *Server) ShowAboutPage(w http.ResponseWriter, r *http.Request) {
    view.ShowAboutPage(w)
}

// TODO Implement call for Slack
