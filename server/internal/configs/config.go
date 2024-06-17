package configs

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"github.com/nitzanpap/url-shortener/server/pkg/colors"
	"github.com/nitzanpap/url-shortener/server/pkg/utils"
)

// LoadConfig loads the configuration from the environment variables
func LoadConfig() *Config {
	// Load configuration from file or any other source

	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	config := &Config{
		Port:         getEnvAsInt("PORT"),
		Database:     getDatabaseConfig(),
		Environment:  getEnvironment(),
		ClientOrigin: os.Getenv("CLIENT_ORIGIN"),
	}

	// if config.Environment is not one of the predefined values, throw an error
	validateEnvironmentVar(config)

	if config.Environment == Development {
		printOutConfig(config)
	}

	v, values := extractConfigFields(config)
	isInvalidConfig, errStringArr := utils.DoesContainEmptyStrings(values, v)

	if isInvalidConfig {
		log.Fatalf(colors.Error("Error loading configuration - Missing values in: %s\n"), strings.Join(errStringArr, ", "))
	}

	return config
}

func getEnvAsInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Fatalf("Error parsing %s: %s", key, err)
	}
	return value
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     getEnvAsInt("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Name:     os.Getenv("DB_NAME"),
		DB_URL:   utils.BuildPostgresqlDbURL(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME")),
	}
}

func getEnvironment() Environment {
	environment := Environment(os.Getenv("ENV"))
	for _, env := range []Environment{Development, Production} {
		if environment == env {
			return environment
		}
	}
	log.Fatalf(colors.Error("Error loading configuration - Invalid environment value: %s\n"), environment)
	return ""
}

func printOutConfig(config *Config) {
	configPrettyJsonStr, err := utils.PrettyStruct(*config)
	if err != nil {
		log.Fatalf(colors.Error("Error pretty printing config: %v"), err)
	}
	log.Printf(colors.Info("Config: %s\n"), configPrettyJsonStr)
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
