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
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			isValidOrigin := strings.Contains(config.ClientOrigin, origin) || strings.HasPrefix(origin, "http://localhost:")
			if !isValidOrigin {
				log.Printf(colors.Error("Invalid origin: %s"), origin)
			}
			return isValidOrigin
		},
		MaxAge: 12 * time.Hour,
	}))
	routes.InitializeRoutes(r)
	return r
}

func setGinMode(config *configs.Config) {
	switch config.Environment {
	case configs.Development:
		gin.SetMode(gin.DebugMode)
	case configs.Production:
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	config := configs.LoadConfig()

	// Set Gin to production mode according to the environment in a switch statement
	setGinMode(config)

	// Create a Gin router instance
	router := setupRouter()

	// Starting the server
	log.Printf(colors.Success("Starting server on: http://localhost:%d\n"), config.Port)
	if err := router.Run(":" + strconv.Itoa(config.Port)); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
