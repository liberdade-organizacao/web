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
    http.HandleFunc("/contato", server.ShowContactPage)
    http.HandleFunc("/suporte", server.ShowSupportPage)
    http.HandleFunc("/blog", server.DisplayBlog)
    http.HandleFunc("/blog/post", server.DisplayPost)
  	http.HandleFunc("/api/posts", server.GivePosts)

    return server
}

func (server *Server) Serve() {
    http.ListenAndServe(server.Port, nil)
}

func (server *Server) ShowIndexPage(w http.ResponseWriter, r *http.Request) {
    posts := model.GetPosts(0)[0:4]
    view.ShowIndex(w, posts)
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

func (server *Server) ShowSupportPage(w http.ResponseWriter, r *http.Request) {
    view.ShowSupport(w)
}

func (server *Server) ShowContactPage(w http.ResponseWriter, r *http.Request) {
    view.ShowContactUs(w)
}

// Sends a message to the people in Liberdade
// TODO Implement me!
// func (server *Server) ContactUs(w http.ResponseWriter, r *http.Request) {
//     view.ShowContactUs(w)
// }

// Provides posts for the outside world via JSON payload
func (server *Server) GivePosts(w http.ResponseWriter, r *http.Request) {
	offset := model.GetOffset(r)
    posts := model.GetPosts(offset)
  	// TODO Set response type to JSON
    view.ProvidePosts(w, posts)
}
