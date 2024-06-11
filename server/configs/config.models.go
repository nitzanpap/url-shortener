package configs

type Config struct {
	Port     string
	Database DatabaseConfig
	Environment Environment
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)
