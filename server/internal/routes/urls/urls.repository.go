package urls

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	dbconfigs "github.com/nitzanpap/url-shortener/server/internal/configs/dbConfigs"
)

func saveUrlInDb(url string, obfuscatedShortenedUrl string, db *pgx.Conn) error {
	_, err := db.Exec(context.Background(), dbconfigs.PreparedStatements.CreateUrlRow, url, obfuscatedShortenedUrl, nil)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr != nil {
			// 23505 means the url already exists in the database, so we can ignore this error
			if pgErr.Code == "23505" {
				return nil
			}
			return err
		}
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
