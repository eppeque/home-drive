package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/eppeque/home-drive/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := createRouter()

	r.Run(":4840")
}

//go:embed app/*
var content embed.FS

func createRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	content, _ := fs.Sub(content, "app")

	r.StaticFS("/app", http.FS(content))
	r.GET("/hello", handlers.HandleHello)

	return r
}
