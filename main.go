package main

import "github.com/liberdade-organizacao/web/controller"

func main() {
    server := controller.NewServer()
    server.Serve()
}
