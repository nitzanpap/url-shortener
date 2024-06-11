package configs

import (
	"log"
	"os"
	"reflect"
	"strconv"

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
		log.Fatalf("Error loading .env file: %s", err)
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

	v, values := extractConfigFields(config)

	log.Printf(colors.Info("Loaded configuration: %#v\n"), values)
	isInvalidConfig, errString := doesContainEmptyValues(values, v)

	if isInvalidConfig {
		log.Fatalf(colors.Error("Error loading configuration: missing values in %s\n"), errString)
	}

	if config.Environment == Development {
		configPrettyJsonStr, err := utils.PrettyStruct(*config)
		if err != nil {
			log.Fatalf(colors.Error("Error pretty printing config: %v"), err)
		}
		log.Printf(colors.Info("Config: %s\n"), configPrettyJsonStr)
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
	for i := 0; i < len(values); i++ {
		if values[i] == "" {
			emptyEntry := v.Type().Field(i).Name
			return true, []string{emptyEntry}
		}
		// if the value is a struct, recursively check for empty values
		if reflect.ValueOf(values[i]).Kind() == reflect.Struct {
			innerValues := make([]interface{}, v.Field(i).NumField())
			for j := 0; j < v.Field(i).NumField(); j++ {
				innerValues[j] = v.Field(i).Field(j).Interface()
			}
			innerStructValidity, err := doesContainEmptyValues(innerValues, v.Field(i))
			if innerStructValidity {
				return true, err
			}
		}
	}
	return false, nil
}
