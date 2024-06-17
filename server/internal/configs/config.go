package configs

import (
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"

	dbconfigs "github.com/nitzanpap/url-shortener/server/internal/configs/dbConfigs"
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
		Port:         utils.GetEnvAsInt("PORT"),
		Database:     dbconfigs.GetDatabaseConfig(),
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
