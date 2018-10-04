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
    http.HandleFunc("/blog/post", server.DisplayPost)
  	http.HandleFunc("/api/posts", server.GivePosts)

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

func (server *Server) DisplayPost(w http.ResponseWriter, r *http.Request) {
    postId := model.GetPostId(r)
    post := model.GetPost(postId)
    view.ShowPost(w, post)
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

    } else {
        message = "oops"
    }

    view.ShowIndex(w, message)
}

// Provides posts for the outside world via JSON payload
func (server *Server) GivePosts(w http.ResponseWriter, r *http.Request) {
	offset := model.GetOffset(r)
    posts := model.GetPosts(offset)
  	// TODO Set response type to JSON
    view.ProvidePosts(w, posts)
}
