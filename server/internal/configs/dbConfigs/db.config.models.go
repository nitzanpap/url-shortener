package dbconfigs

type DatabaseConfig struct {
	ConnectionType DbConnectionType
	DirectURL      string
	Host           string
	Port           int
	Username       string
	Password       string
	Name           string
	DB_URL         string
}

type DbConnectionType string

const (
	SingleConnection DbConnectionType = "direct"
	PoolConnection   DbConnectionType = "pool"
)

type preparedStatementsStruct struct {
	CreateUserRow, GetUserRow, CreateUrlRow, GetUrlRow, GetUrlsByUserId string
}
