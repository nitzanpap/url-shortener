package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nitzanpap/url-shortener/configs"
	"github.com/nitzanpap/url-shortener/internal/routes"
	"github.com/nitzanpap/url-shortener/pkg/colors"
)

func setupRouter() *gin.Engine {
	config := configs.LoadConfig()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == config.ClientOrigin || strings.HasPrefix(origin, "http://localhost:")
		},
		MaxAge: 12 * time.Hour,
	}))
	routes.InitializeRoutes(r)
	return r
}

func main() {
	config := configs.LoadConfig()

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Create a Gin router instance
	router := setupRouter()

	// Starting the server
	log.Printf(colors.Success("Starting server on: http://localhost:%d\n"), config.Port)
	if err := router.Run(":" + strconv.Itoa(config.Port)); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
