package urls

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/nitzanpap/url-shortener/server/pkg/colors"
)

// this route is: POST /h/:obfuscatedShortenedUrl, and it returns the original URL
func ShortUrlHandler(r *gin.RouterGroup, db *pgx.Conn) {
	r.GET("/:obfuscatedShortenedUrl", func(c *gin.Context) {
		obfuscatedShortenedUrl := c.Param("obfuscatedShortenedUrl")
		actualUrl, err := getUrl(obfuscatedShortenedUrl, db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get URL"})
			return
		}
		if !isUrlValid(actualUrl) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid URL"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"originalUrl": actualUrl})
	})
}

// This route is: GET /api/v1/url/:obfuscatedShortenedUrl
func UrlGroupHandler(r *gin.RouterGroup, db *pgx.Conn) {
	url := r.Group("/url")
	{
		// This route receives a POST request with a JSON body that contains a URL, and returns a shortened URL.
		url.POST("/", func(c *gin.Context) {
			var request struct {
				URL string `json:"url"`
			}
			if err := c.ShouldBindJSON(&request); err != nil {
				log.Fatalf(colors.Error("Failed to bind JSON: %s"), err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
				return
			}
			if !isUrlValid(request.URL) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
				return
			}
			obfuscatedShortenedUrl, err := saveUrl(request.URL, db)
			if err != nil {
				log.Fatalf(colors.Error("Failed to save URL: %s"), err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short URL"})
				return
			}

			c.JSON(http.StatusCreated, gin.H{"obfuscatedShortenedUrl": obfuscatedShortenedUrl})
		})
	}
}

func isUrlValid(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}
