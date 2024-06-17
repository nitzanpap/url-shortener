package configs

type Config struct {
	Port        int
	Database    DatabaseConfig
	Environment Environment
	ClientOrigin string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	DB_URL   string
}

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

type DbConnectionType string

const (
	SingleConnection DbConnectionType = "single"
	PoolConnection   DbConnectionType = "pool"
)

type preparedStatementsStruct struct {
	CreateUserRow, GetUserRow, CreateUrlRow, GetUrlRow, GetUrlsByUserId string
}
