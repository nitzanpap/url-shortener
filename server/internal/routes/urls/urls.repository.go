package urls

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	dbconfigs "github.com/nitzanpap/url-shortener/server/internal/configs/dbConfigs"
)

func saveURLInDB(url string, obfuscatedShortenedURL string, db *pgx.Conn) error {
	_, err := db.Exec(context.Background(), dbconfigs.PreparedStatements.CreateURLRow, url, obfuscatedShortenedURL, nil)
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

func getURLFromDB(obfuscatedShortenedURL string, db *pgx.Conn) (string, error) {
	var url string
	err := db.QueryRow(context.Background(), `SELECT original_url FROM urls WHERE obfuscated_shortened_url = $1`, obfuscatedShortenedURL).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}

func getBase62StringFromDB(url string, db *pgx.Conn) (string, error) {
	var base62String string
	err := db.QueryRow(context.Background(), `SELECT obfuscated_shortened_url FROM urls WHERE original_url = $1`, url).Scan(&base62String)
	if err != nil {
		return "", err
	}
	return base62String, nil
}
