package urls

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UrlsGroupHandler(r *gin.RouterGroup) {
	urls := r.Group("/urls")
	{
		// This route receives a POST request with a JSON body that contains a URL, and returns a shortened URL.
		urls.POST("/", func(c *gin.Context) {
			var request struct {
				URL string `json:"url"`
			}
			if err := c.ShouldBindJSON(&request); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
				return
			}
			shortURL, err := generateShortURL(request.URL)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short URL"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
		})
	}
}
