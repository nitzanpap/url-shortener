package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadConfig() (*Config, error) {
	// Load configuration from file or any other source

	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Getting and using a value from .env
	Port := os.Getenv("PORT")
	DbHost := os.Getenv("DB_HOST")
	DBPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	DbUser := os.Getenv("DB_USER")
	DbPass := os.Getenv("DB_PASS")
	DbName := os.Getenv("DB_NAME")

	if err == nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	config := &Config{
		Port: Port,
		Database: DatabaseConfig{
			Host:     DbHost,
			Port:     DBPort,
			Username: DbUser,
			Password: DbPass,
			Name:     DbName,
		},
	}

	return config, nil
}

// usage examples:

// config, err := configs.LoadConfig()
// if err != nil {
// 	log.Fatalf("Error loading configuration: %v", err)
// }
// fmt.Printf("Config: %#v", config)

// fmt.Printf("Port: %s", config.Port)

// fmt.Printf("Database host: %s", config.Database.Host)

// fmt.Printf("Database port: %d", config.Database.Port)

// fmt.Printf("Database username: %s", config.Database.Username)

// fmt.Printf("Database password: %s", config.Database.Password)

// fmt.Printf("Database name: %s", config.Database.Name)

// Output:
// Config: &configs.Config{Port:"8080", Database:configs.DatabaseConfig{Host:"localhost", Port:543

// Port: 8080
// Database host: localhost
// Database port: 5432
// Database username: myuser
// Database password: mypassword
// Database name: mydatabase

// Note: The output may vary depending on the configuration values
