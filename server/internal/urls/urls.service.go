package urls

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/jackc/pgx/v5"
)

func saveUrl(url string, db *pgx.Conn) (string, error) {
	// generate a hashedUrl for the URL
	hashedUrl := hashUrl(url)

	// save the URL and the hash in the database
	err := saveUrlInDb(url, hashedUrl, db)
	if err != nil {
		return "", err
	}

	// return the shortened URL
	return hashedUrl, nil
}

func getUrl(hash string, db *pgx.Conn) (string, error) {
	// get the URL from the database using the hash
	url, err := getUrlFromDb(hash, db)
	if err != nil {
		return "", err
	}

	// return the URL
	return url, nil
}

func hashUrl(url string) string {
	// generate a hash for the URL
	// return the hash
	hash := sha256.Sum256([]byte(url))
	return hex.EncodeToString(hash[:])
}
