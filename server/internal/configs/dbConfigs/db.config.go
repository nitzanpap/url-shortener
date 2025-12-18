package dbconfigs

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nitzanpap/url-shortener/server/pkg/colors"
	"github.com/nitzanpap/url-shortener/server/pkg/utils"
)

func GetDatabaseConfig() DatabaseConfig {
	connectionType := DbConnectionType(os.Getenv("DB_CONNECTION_TYPE"))
	if connectionType == "" {
		connectionType = SingleConnection // Default to single connection
	}

	return DatabaseConfig{
		ConnectionType: connectionType,
		DirectURL:      os.Getenv("DB_DIRECT_URL"),
		Host:           os.Getenv("DB_HOST"),
		Port:           utils.GetEnvAsInt("DB_PORT"),
		Username:       os.Getenv("DB_USER"),
		Password:       os.Getenv("DB_PASS"),
		Name:           os.Getenv("DB_NAME"),
		DB_URL:         utils.BuildPostgresqlDbURL(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME")),
	}
}

// ConnectToDB connects to the database based on the provided configuration
func ConnectToDB(config DatabaseConfig) (*pgx.Conn, error) {
	var conn *pgx.Conn
	var err error

	// Try direct connection first if configured
	if config.ConnectionType == SingleConnection && config.DirectURL != "" {
		// Force IPv4 by adding ?prefer_socket=true to the connection string
		directURL := config.DirectURL
		if !strings.Contains(directURL, "?") {
			directURL += "?prefer_socket=true"
		} else {
			directURL += "&prefer_socket=true"
		}

		log.Print(colors.Info("Attempting direct connection to database...\n"))
		conn, err = pgx.Connect(context.Background(), directURL)
		if err == nil {
			log.Print(colors.Success("Successfully connected using direct connection\n"))
			return conn, nil
		}

		log.Printf("%s%v\n", colors.Warning("Direct connection failed: "), err)
		log.Print(colors.Info("Falling back to pooler connection...\n"))
	}

	// Use pooler connection if direct failed or wasn't configured
	log.Print(colors.Info("Attempting to connect via connection pooler...\n"))
	conn, err = pgx.Connect(context.Background(), config.DB_URL)
	if err != nil {
		return nil, err
	}

	log.Print(colors.Success("Successfully connected using pooler connection\n"))
	return conn, nil
}

func ConnectToDBPool(dbURL string) (*pgxpool.Pool, error) {
	// Connect to the database via pgx for pool connection
	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf(colors.Error("Unable to create connection pool: %v\n"), err)
	}
	return dbPool, nil
}

func InitDB(db *pgx.Conn) {
	createDbTables(db)
	enableRLS(db)
	createPreparedStatements(db)
}

func createDbTables(db *pgx.Conn) {
	// Create the Users table if it does not exist
	_, err := db.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		hashed_password TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UNIQUE (username)
		);`)
	if err != nil {
		log.Fatalf(colors.Error("Unable to create users table: %v\n"), err)
	}

	// Create the URLs table if it does not exist.
	_, err = db.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS urls (
        id SERIAL PRIMARY KEY,
        user_id INTEGER,
        original_url TEXT NOT NULL,
        obfuscated_shortened_url TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(user_id),
        UNIQUE (original_url),
        UNIQUE (obfuscated_shortened_url)
    );
	`)
	if err != nil {
		log.Fatalf(colors.Error("Unable to create urls table: %v\n"), err)
	}
}

func enableRLS(db *pgx.Conn) {
	// Enable RLS on the Users table
	_, err := db.Exec(context.Background(), "ALTER TABLE users ENABLE ROW LEVEL SECURITY;")
	if err != nil {
		log.Fatalf(colors.Error("Unable to enable RLS on users table: %v\n"), err)
	}

	// Enable RLS on the URLs table
	_, err = db.Exec(context.Background(), "ALTER TABLE urls ENABLE ROW LEVEL SECURITY;")
	if err != nil {
		log.Fatalf(colors.Error("Unable to enable RLS on urls table: %v\n"), err)
	}
}

var PreparedStatements = preparedStatementsStruct{
	CreateUserRow:   "createUserRow",
	GetUserRow:      "getUserRow",
	CreateUrlRow:    "createUrlRow",
	GetUrlRow:       "getUrlRow",
	GetUrlsByUserId: "getUrlsByUserId",
}

func createPreparedStatements(db *pgx.Conn) {
	preparedStatements := map[string]string{
		PreparedStatements.CreateUserRow:   `INSERT INTO users (username, hashed_password) VALUES ($1, $2)`,
		PreparedStatements.GetUserRow:      `SELECT user_id, username, hashed_password FROM users WHERE username = $1`,
		PreparedStatements.CreateUrlRow:    `INSERT INTO urls (original_url, obfuscated_shortened_url, user_id) VALUES ($1, $2, $3)`,
		PreparedStatements.GetUrlRow:       `SELECT id, user_id, original_url, obfuscated_shortened_url FROM urls WHERE obfuscated_shortened_url = $1`,
		PreparedStatements.GetUrlsByUserId: `SELECT id, user_id, original_url, obfuscated_shortened_url FROM urls WHERE user_id = $1`,
	}

	_, err := db.Exec(context.Background(), "DEALLOCATE ALL")
	if err != nil {
		log.Fatalf(colors.Error("Failed to deallocate prepared statements: %s\n"), err)
	}

	for stmtName, stmtQuery := range preparedStatements {
		_, err = db.Prepare(context.Background(), stmtName, stmtQuery)
		if err != nil {
			log.Fatalf(colors.Error("Failed to create prepared statement %s: %s\n"), stmtName, err)
		}
	}
}
