package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//go:embed index.html
var indexHTML embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		data, _ := indexHTML.ReadFile("index.html")
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Printf("guitar-fingerboard listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
