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

	DBPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err == nil {
		log.Fatalf("Error parsing DB_PORT: %s", err)
	}

	config := &Config{
		Port: os.Getenv("PORT"),
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     DBPort,
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Name:     os.Getenv("DB_NAME"),
		},
		Environment: Environment(os.Getenv("ENV")),
	}

	if config.Environment == Development {
		configPrettyJsonStr, err := utils.PrettyStruct(*config)
		if err != nil {
			log.Fatalf(colors.Error("Error pretty printing config: %v"), err)
		}
		log.Printf(colors.Info("Config: %s\n"), configPrettyJsonStr)
	}

	v, values := extractConfigFields(config)
	isInvalidConfig, errStringArr := doesContainEmptyValues(values, v)

	if isInvalidConfig {
		log.Fatalf(colors.Error("Error loading configuration - Missing values in: %s\n"), strings.Join(errStringArr, ", "))
	}

	return config, nil
}

func extractConfigFields(config *Config) (reflect.Value, []interface{}) {
	v := reflect.ValueOf(*config)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return v, values
}

func doesContainEmptyValues(values []interface{}, v reflect.Value) (bool, []string) {
	var emptyFields []string
	for i := 0; i < len(values); i++ {
		if values[i] == "" {
			emptyEntry := v.Type().Field(i).Name
			emptyFields = append(emptyFields, emptyEntry)
		} else if reflect.ValueOf(values[i]).Kind() == reflect.Struct {
			// if the value is a struct, recursively check for empty values
			innerValues := make([]interface{}, v.Field(i).NumField())
			for j := 0; j < v.Field(i).NumField(); j++ {
				innerValues[j] = v.Field(i).Field(j).Interface()
			}
			isInvalidConfig, errStringArr := doesContainEmptyValues(innerValues, v.Field(i))
			if isInvalidConfig {
				emptyFields = append(emptyFields, errStringArr...)
			}
		}
	}
	return len(emptyFields) > 0, emptyFields
}
