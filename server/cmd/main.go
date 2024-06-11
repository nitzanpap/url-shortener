package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nitzanpap/url-shortener/configs"
	"github.com/nitzanpap/url-shortener/internal/routes"
	"github.com/nitzanpap/url-shortener/pkg/colors"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.InitializeRoutes(r)
	return r
}

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf(colors.Error("Error loading configuration: %v"), err)
	}

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
