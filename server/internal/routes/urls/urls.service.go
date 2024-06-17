package urls

import (
	"crypto/sha256"

	"github.com/jackc/pgx/v5"
	"github.com/nitzanpap/url-shortener/server/pkg/utils"
)

func saveUrl(url string, db *pgx.Conn) (string, error) {
	obfuscatedShortenedUrl := shortenAndObfuscateStringUniquely(url)

	// save the URL and the obfuscatedShortenedUrl in the database
	err := saveUrlInDb(url, obfuscatedShortenedUrl, db)
	if err != nil {
		return "", err
	}
	return obfuscatedShortenedUrl, nil
}

func getUrl(obfuscatedShortenedUrl string, db *pgx.Conn) (string, error) {
	// get the URL from the database using the obfuscatedShortenedUrl
	url, err := getUrlFromDb(obfuscatedShortenedUrl, db)
	if err != nil {
		return "", err
	}
	return url, nil
}

func shortenAndObfuscateStringUniquely(url string) string {
	hash := sha256.Sum256([]byte(url))
	base62String := utils.Base62Encode(hash[:])
	return base62String
}
