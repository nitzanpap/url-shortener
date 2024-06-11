package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitzanpap/url-shortener/internal/handlers"
)

func InitializeRoutes(r *gin.Engine) {
	var db = make(map[string]string)

	r.GET("/", handlers.HomeHandler)
	r.GET("/about", handlers.AboutHandler)

	v1 := r.Group("/v1")
	{

		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok", "version": "v1"})
		})

		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		// Get user value
		v1.GET("/user/:name", func(c *gin.Context) {
			user := c.Params.ByName("name")
			value, ok := db[user]
			if ok {
				c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
			} else {
				c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
			}
		})

	}
}
