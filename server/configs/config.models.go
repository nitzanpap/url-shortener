package configs

type Config struct {
	Port     string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}
