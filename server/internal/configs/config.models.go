package configs

import dbconfigs "github.com/nitzanpap/url-shortener/server/internal/configs/dbConfigs"

type Config struct {
	Port         int
	Database     dbconfigs.DatabaseConfig
	Environment  Environment
	ClientOrigin string
}
type Environment string

const (
	Local       Environment = "local"
	Development Environment = "development"
	Production  Environment = "production"
)
