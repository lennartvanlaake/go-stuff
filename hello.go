package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
	"github.com/htmx-poc/renderer"
	. "github.com/htmx-poc/templates"
	. "github.com/htmx-poc/types"
)

func getUsers(db *sql.DB) []User {
	var users []User
	rows, err := db.Query("SELECT * from users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.Name)
		users = append(users, user)
	}
	return users
}

// https://gin-gonic.com/docs/examples/html-rendering
func main() {
	router := gin.Default()
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
		c.HTML(http.StatusOK, "", Admin(getUsers(db)))
	})

	router.POST("/admin", func(c *gin.Context) {
		c.Request.ParseForm()
		name := c.PostForm("fname")
		db.Exec("INSERT INTO users VALUES (?)", name)
		c.HTML(http.StatusOK, "", UserList(getUsers(db)))
	})

	router.DELETE("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		db.Exec("DELETE FROM users WHERE name = (?)", name)
		c.HTML(http.StatusOK, "", UserList(getUsers(db)))
	})

	router.Run()
}
