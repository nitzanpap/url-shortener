package dbconfigs

type DatabaseConfig struct {
	ConnectionType DBConnectionType
	DirectURL      string `json:"-"`
	Host           string
	Port           int
	Username       string
	Password       string `json:"-"`
	Name           string
	DBURL          string `json:"-"`
}

type DBConnectionType string

const (
	SingleConnection DBConnectionType = "direct"
	PoolConnection   DBConnectionType = "pool"
)

type preparedStatementsStruct struct {
	CreateURLRow, GetURLRow, GetURLsByUserID string
}
