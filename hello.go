package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

// https://gin-gonic.com/docs/examples/html-rendering
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	router.GET("/index", func(c *gin.Context) {
		var version string
		row := db.QueryRow("select sqlite_version()")
		row.Scan(&version)
		log.Print(version)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"name": version,
		})
	})
	router.Run()
}
