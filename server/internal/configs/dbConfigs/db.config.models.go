package dbconfigs

type DatabaseConfig struct {
	ConnectionType DBConnectionType
	DirectURL      string
	Host           string
	Port           int
	Username       string
	Password       string
	Name           string
	DBURL          string
}

type DBConnectionType string

const (
	SingleConnection DBConnectionType = "direct"
	PoolConnection   DBConnectionType = "pool"
)

type preparedStatementsStruct struct {
	CreateUserRow, GetUserRow, CreateURLRow, GetURLRow, GetURLsByUserID string
}
