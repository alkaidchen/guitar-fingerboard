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

	const iconSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><rect width="100" height="100" rx="20" fill="#1a1a2e"/><text x="50" y="50" font-size="70" text-anchor="middle" dominant-baseline="central">🎸</text></svg>`

	r.GET("/icon.svg", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/svg+xml; charset=utf-8", []byte(iconSVG))
	})

	r.GET("/manifest.json", func(c *gin.Context) {
		manifest := `{
  "name": "吉他指板学习",
  "short_name": "指板学习",
  "start_url": "/",
  "display": "fullscreen",
  "orientation": "landscape",
  "background_color": "#1a1a2e",
  "theme_color": "#1a1a2e",
  "icons": [
    {
      "src": "/icon.svg",
      "sizes": "any",
      "type": "image/svg+xml",
      "purpose": "any"
    }
  ]
}`
		c.Data(http.StatusOK, "application/manifest+json; charset=utf-8", []byte(manifest))
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Printf("guitar-fingerboard listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
