package main

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/nitzanpap/url-shortener/server/internal/configs"
	"github.com/nitzanpap/url-shortener/server/internal/routes"
	"github.com/nitzanpap/url-shortener/server/pkg/colors"
	"github.com/nitzanpap/url-shortener/server/pkg/utils"
)

func main() {
	config := configs.LoadConfig()

	gin.SetMode(configs.GetGinMode(config))

	db, err := configs.ConnectToDB(config.Database.DB_URL)
	if err != nil {
		log.Fatalf(colors.Error("could not connect to database: %s\n"), err)
	}
	defer db.Close(context.Background())

	// Test the connection to the database
	utils.TestDBConnection(db)

	configs.InitDB(db)
	log.Print(colors.Success("Successfully initialized database\n"))

	// Create a Gin router instance
	router := configs.SetupGinServer(config)
	routes.InitializeRoutes(router, db)

	// Starting the server
	log.Printf(colors.Success("Starting server on: http://localhost:%d\n"), config.Port)
	if err := router.Run(":" + strconv.Itoa(config.Port)); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
