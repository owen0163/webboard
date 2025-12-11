package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	db, err := sql.Open("sqlite3", "./webboard.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r.StaticFile("/", "./webboard.html")
	r.StaticFile("/webboard.html", "./webboard.html")
	// GET
	r.GET("/api/boards", func(c *gin.Context) {
		rows, err := db.Query("SELECT name, massage FROM webboard")
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{
				"Massage": err.Error(),
			})
			return
		}
		var webboard []board
		for rows.Next() {
			var name, msg string
			if err := rows.Scan(&name, &msg); err != nil {
				log.Println(err)
			}
			webboard = append(webboard, board{
				Name:    name,
				Message: msg,
			})
		}
		c.JSON(http.StatusOK, webboard)
	})
	// POST
	r.POST("/api/boards", func(c *gin.Context) {
		var b board
		if err := c.BindJSON(&b); err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{
				"Massage": err.Error(),
			})
			return
		}
		db.Exec("INSERT INTO webboard(name, massage) VALUES(?,?)", b.Name, b.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{
				"Massage": err.Error(),
			})
			return
		}
		rows, err := db.Query("SELECT name, massage FROM webboard")
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{
				"Massage": err.Error(),
			})
			return
		}
		var webboard []board
		for rows.Next() {
			var name, msg string
			if err := rows.Scan(&name, &msg); err != nil {
				log.Println(err)
			}
			webboard = append(webboard, board{
				Name:    name,
				Message: msg,
			})
		}
		c.JSON(http.StatusOK, webboard)

	})
	fmt.Println("listening and serving on :", os.Getenv("PORT"))
	r.Run()
	fmt.Println("bye")
}

var boards = []board{}

type board struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
