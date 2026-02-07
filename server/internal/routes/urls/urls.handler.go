package urls

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/nitzanpap/url-shortener/server/pkg/colors"
)

// ShortURLHandler handles the route POST /h/:obfuscatedShortenedUrl, and returns the original URL.
func ShortURLHandler(r *gin.RouterGroup, db *pgx.Conn) {
	r.GET("/:obfuscatedShortenedUrl", func(c *gin.Context) {
		obfuscatedShortenedURL := c.Param("obfuscatedShortenedUrl")
		actualURL, err := getURL(obfuscatedShortenedURL, db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get URL"})
			return
		}
		if !isURLValid(actualURL) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid URL"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"originalUrl": actualURL})
	})
}

// URLGroupHandler handles the route GET /api/v1/url/:obfuscatedShortenedUrl.
func URLGroupHandler(r *gin.RouterGroup, db *pgx.Conn) {
	url := r.Group("/url")
	{
		// This route receives a POST request with a JSON body that contains a URL, and returns a shortened URL.
		url.POST("/", func(c *gin.Context) {
			var request struct {
				URL string `json:"url"`
			}
			if err := c.ShouldBindJSON(&request); err != nil {
				log.Printf(colors.Error("Failed to bind JSON: %s"), err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
				return
			}
			if !isURLValid(request.URL) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
				return
			}
			obfuscatedShortenedURL, err := saveURL(request.URL, db)
			if err != nil {
				log.Printf(colors.Error("Failed to save URL: %s"), err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short URL"})
				return
			}

			c.JSON(http.StatusCreated, gin.H{"obfuscatedShortenedUrl": obfuscatedShortenedURL})
		})
	}
}

func isURLValid(rawURL string) bool {
	trimmed := strings.TrimSpace(rawURL)
	if trimmed == "" {
		return false
	}
	parsed, err := url.ParseRequestURI(trimmed)
	if err != nil {
		return false
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return false
	}
	if parsed.Host == "" {
		return false
	}
	return true
}
