package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitzanpap/url-shortener/configs"
	"github.com/nitzanpap/url-shortener/pkg/colors"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	v1 := r.Group("/v1")
	{

		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok", "version": "v1"})
		})

		// Ping test
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
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

		// Authorized group (uses gin.BasicAuth() middleware)
		// Same than:
		// authorized := v1.Group("/")
		// authorized.Use(gin.BasicAuth(gin.Credentials{
		//	  "foo":  "bar",
		//	  "manu": "123",
		//}))
		authorized := v1.Group("/", gin.BasicAuth(gin.Accounts{
			"foo":  "bar", // user:foo password:bar
			"manu": "123", // user:manu password:123
		}))

		/* example curl for /admin with basicauth header
		   Zm9vOmJhcg== is base64("foo:bar")

			curl -X POST \
		  	http://localhost:8080/admin \
		  	-H 'authorization: Basic Zm9vOmJhcg==' \
		  	-H 'content-type: application/json' \
		  	-d '{"value":"bar"}'
		*/
		authorized.POST("admin", func(c *gin.Context) {
			user := c.MustGet(gin.AuthUserKey).(string)

			// Parse JSON
			var json struct {
				Value string `json:"value" binding:"required"`
			}

			if c.Bind(&json) == nil {
				db[user] = json.Value
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
			}
		})

		return r
	}
}

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf(colors.Error("Error loading configuration: %v"), err)
	}

	r := setupRouter()
	fmt.Printf(colors.Success("Server running on: http://localhost:%s\n"), config.Port)
	// Listen and Server in 0.0.0.0:${port}
	r.Run(":" + config.Port)
}
