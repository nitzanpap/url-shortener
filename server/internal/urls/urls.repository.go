package urls

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// This is the repository layer for the URLs service. It is responsible for interacting with the database.

func saveUrlInDb(url string, hashedUrl string, db *pgx.Conn) error {
	_, err := db.Exec(context.Background(), `INSERT INTO urls (original_url, hash) VALUES ($1, $2)`, url, hashedUrl)
	if err != nil {
		return err
	}
	return nil
}

func getUrlFromDb(hash string, db *pgx.Conn) (string, error) {
	var url string
	err := db.QueryRow(context.Background(), `SELECT original_url FROM urls WHERE hash = $1`, hash).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
