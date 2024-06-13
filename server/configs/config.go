package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/nitzanpap/url-shortener/pkg/colors"
	"github.com/nitzanpap/url-shortener/pkg/utils"
)

// LoadConfig loads the configuration from the environment variables
func LoadConfig() *Config {
	// Load configuration from file or any other source

	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	Port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Error parsing PORT: %s", err)
	}
	DB_HOST := os.Getenv("DB_HOST")

	DB_Port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error parsing DB_PORT: %s", err)
	}
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_URL := buildDbURL(DB_HOST, os.Getenv("DB_PORT"), DB_USER, DB_PASS, DB_NAME)
	Environment := Environment(os.Getenv("ENV"))
	ClientOrigin := os.Getenv("CLIENT_ORIGIN")

	config := &Config{
		Port: Port,
		Database: DatabaseConfig{
			Host:     DB_HOST,
			Port:     DB_Port,
			Username: DB_USER,
			Password: DB_PASS,
			Name:     DB_NAME,
			DB_URL:   DB_URL,
		},
		Environment:  Environment,
		ClientOrigin: ClientOrigin,
	}

	// if config.Environment is not one of the predefined values, throw an error
	validateEnvironmentVar(config)

	if config.Environment == Development {
		configPrettyJsonStr, err := utils.PrettyStruct(*config)
		if err != nil {
			log.Fatalf(colors.Error("Error pretty printing config: %v"), err)
		}
		log.Printf(colors.Info("Config: %s\n"), configPrettyJsonStr)
	}

	v, values := extractConfigFields(config)
	isInvalidConfig, errStringArr := utils.DoesContainEmptyStrings(values, v)

	if isInvalidConfig {
		log.Fatalf(colors.Error("Error loading configuration - Missing values in: %s\n"), strings.Join(errStringArr, ", "))
	}

	return config
}

// Define the buildDbURL function
func buildDbURL(host, port, user, pass, name string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, name)
}

// returns either a (pgx.Conn, error) or a (pgxpool.Pool, error)
func ConnectToDB(dbURL string) (*pgx.Conn, error) {
	// Connect to the database via pgx for single connection
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func ConnectToDBPool(dbURL string) (*pgxpool.Pool, error) {
	// Connect to the database via pgx for pool connection
	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf(colors.Error("Unable to create connection pool: %v\n"), err)
	}
	return dbPool, nil
}

func validateEnvironmentVar(config *Config) {
	for _, env := range []Environment{Development, Production} {
		if config.Environment == env {
			return
		}
	}
	log.Fatalf(colors.Error("Error loading configuration - Invalid environment value: %s\n"), config.Environment)
}

func extractConfigFields(config *Config) (reflect.Value, []interface{}) {
	v := reflect.ValueOf(*config)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return v, values
}
