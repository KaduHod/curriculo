package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
const PUBLIC_DIR = "./public"
const PAGES_DIR = PUBLIC_DIR+"/pages"
const JS_FILES = PUBLIC_DIR+"/js"

func main() {
	server := gin.Default()
	server.LoadHTMLGlob(PAGES_DIR +"/*")
	server.Static("/public", PUBLIC_DIR)
	fmt.Println(PUBLIC_DIR + "/index.html")
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	server.Run(":3333")
}

