package configs

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/nitzanpap/url-shortener/pkg/colors"
	"github.com/nitzanpap/url-shortener/pkg/utils"
)

// LoadConfig loads the configuration from the environment variables
func LoadConfig() (*Config, error) {
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

	DBPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error parsing DB_PORT: %s", err)
	}

	config := &Config{
		Port: Port,
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     DBPort,
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Name:     os.Getenv("DB_NAME"),
		},
		Environment: Environment(os.Getenv("ENV")),
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

	return config, nil
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
