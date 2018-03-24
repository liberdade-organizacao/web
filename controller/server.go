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
    http.HandleFunc("/", server.SayHello)
    http.HandleFunc("/office", server.TalkAboutOffice)
    http.HandleFunc("/garage", server.TalkAboutGarage)
    http.HandleFunc("/office/gtd", server.TalkAboutGtd)
    http.HandleFunc("/api/mail", server.SendMail)

    return server
}

// Puts the webserver to, well, serve.
func (server *Server) Serve() {
    http.ListenAndServe(server.Port, nil)
}

// Displays the main page
func (server *Server) SayHello(w http.ResponseWriter, r *http.Request) {
    view.SayHello(w)
}

// Displays the office page
func (server *Server) TalkAboutOffice(w http.ResponseWriter, r *http.Request) {
    view.TalkAboutOffice(w)
}

// Displays the garage page
func (server *Server) TalkAboutGarage(w http.ResponseWriter, r *http.Request) {
    view.TalkAboutGarage(w)
}

// Starts the GTD presentation
func (server *Server) TalkAboutGtd(w http.ResponseWriter, r *http.Request) {
    view.TalkAboutGtd(w)
}

// Sends an e-mail to myself
func (server *Server) SendMail(w http.ResponseWriter, r *http.Request) {
    to := r.FormValue("to")
    msg := r.FormValue("msg")
    oops := model.SendSimpleMail(to, msg)
    view.DisplayError(w, oops)
}
