package main

import (
	"database/sql"
	"htmx-server/renderer"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

type User struct {
	Name string
}

// https://gin-gonic.com/docs/examples/html-rendering
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	db, err := sql.Open("sqlite", ":memory:")
	ginHtmlRenderer := router.HTMLRender
	router.HTMLRender = &renderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	if err != nil {
		log.Fatal(err)
	}

	db.Exec("CREATE TABLE users (name text)")

	router.GET("/index", func(c *gin.Context) {
		var version string
		row := db.QueryRow("select sqlite_version()")
		row.Scan(&version)
		log.Print(version)
		c.HTML(http.StatusOK, "", Home())
	})

	router.GET("/admin", func(c *gin.Context) {
		var users []User
		rows, _ := db.Query("SELECT * from users")
		for rows.Next() {
			var user User
			rows.Scan(&user.Name)
			log.Print(user.Name)
			users = append(users, user)
		}
		c.HTML(http.StatusOK, "admin.tmpl", gin.H{
			"users": users,
		})
	})

	router.POST("/admin", func(c *gin.Context) {
		name := c.PostForm("fname")
		log.Print("Processing the post")
		log.Print(name)
		db.Exec("INSERT INTO users VALUES (?)", name)
		c.Redirect(http.StatusFound, "/admin")

	})

	router.Run()
}
