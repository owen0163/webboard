package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	r.StaticFile("/", "./webboard.html")
	r.StaticFile("/webboard.html", "./webboard.html")

	r.GET("/api/boards", func(c *gin.Context) {})
	r.POST("/api/boards", func(c *gin.Context) {})
	fmt.Println("listening and serving on :", os.Getenv("PORT"))
	r.Run()
	fmt.Println("bye")
}

type board struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
