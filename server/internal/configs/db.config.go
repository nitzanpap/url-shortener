package configs

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nitzanpap/url-shortener/server/pkg/colors"
)

// returns either a (pgx.Conn, error) or a (pgxpool.Pool, error)
func ConnectToDB(dbURL string) (*pgx.Conn, error) {
	// Connect to the database via pgx for single connection
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
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
	createPreparedStatements(db)
}

func createDbTables(db *pgx.Conn) {
	// Create the Users table if it does not exist
	_, err := db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS users (
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
	_, err = db.Exec(context.Background(), `
    CREATE TABLE IF NOT EXISTS urls (
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
