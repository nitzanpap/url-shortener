package urls

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func saveUrlInDb(url string, obfuscatedShortenedUrl string, db *pgx.Conn) error {
	_, err := db.Prepare(context.Background(), "createUrlRow", `INSERT INTO urls (original_url, obfuscated_shortened_url) VALUES ($1, $2)`)
	db.Exec(context.Background(), "createUrlRow", url, obfuscatedShortenedUrl)
	if err != nil {
		return err
	}
	return nil
}

func getUrlFromDb(obfuscatedShortenedUrl string, db *pgx.Conn) (string, error) {
	var url string
	err := db.QueryRow(context.Background(), `SELECT original_url FROM urls WHERE obfuscated_shortened_url = $1`, obfuscatedShortenedUrl).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
