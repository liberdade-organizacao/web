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
    http.HandleFunc("/contato", server.ShowAboutPage)
    http.HandleFunc("/contatar", server.ContactUs)
    http.HandleFunc("/suporte", server.ShowSupportPage)

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
func (server *Server) ShowAboutPage(w http.ResponseWriter, r *http.Request) {
    view.ShowAboutPage(w, "")
}

// Displays the garage page
func (server *Server) ShowSupportPage(w http.ResponseWriter, r *http.Request) {
    view.ShowSupportPage(w)
}

// Sends a message to the people in Liberdade
func (server *Server) ContactUs(w http.ResponseWriter, r *http.Request) {
    oops := model.Contact(r)
    message := "ok"

    if oops == nil {

    }

    view.ShowAboutPage(w, message)
}
