package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// https://gin-gonic.com/docs/examples/html-rendering
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"name": "Lennart",
		})
	})
	router.Run()
}
