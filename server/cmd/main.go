package main

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/nitzanpap/url-shortener/server/configs"
	"github.com/nitzanpap/url-shortener/server/internal/routes"
	"github.com/nitzanpap/url-shortener/server/pkg/colors"
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

func GetGinMode(config *configs.Config) (mode string) {
	switch config.Environment {
	case configs.Development:
		return gin.DebugMode
	case configs.Production:
		return gin.ReleaseMode
	default:
		log.Fatalf(colors.Error("Invalid environment: %s"), config.Environment)
		return gin.DebugMode
	}
}

func main() {
	config := configs.LoadConfig()

	gin.SetMode(GetGinMode(config))

	dbPool, err := configs.ConnectToDBPool(config.Database.DB_URL)
	if err != nil {
		log.Fatalf(colors.Error("could not connect to database: %s\n"), err)
	}
	defer dbPool.Close()

	// Test the connection to the database and print the response
	if err := dbPool.Ping(context.Background()); err != nil {
		log.Fatalf(colors.Error("could not ping database: %s\n"), err)
	}
	log.Print(colors.Success("Successfully connected to database\n"))

	configs.InitDB(dbPool)
	log.Print(colors.Success("Successfully initialized database\n"))

	// Create a Gin router instance
	router := setupRouter()

	// Starting the server
	log.Printf(colors.Success("Starting server on: http://localhost:%d\n"), config.Port)
	if err := router.Run(":" + strconv.Itoa(config.Port)); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
