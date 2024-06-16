package urls

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// This is the repository layer for the URLs service. It is responsible for interacting with the database.

func saveUrlInDb(url string, hashedUrl string, dbPool *pgxpool.Pool) error {
	_, err := dbPool.Exec(context.Background(), `INSERT INTO urls (original_url, short_url) VALUES ($1, $2)`, url, hashedUrl)
	if err != nil {
		return err
	}
	return nil
}

func getUrlFromDb(hash string, dbPool *pgxpool.Pool) (string, error) {
	var url string
	err := dbPool.QueryRow(context.Background(), `SELECT original_url FROM urls WHERE short_url = $1`, hash).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
