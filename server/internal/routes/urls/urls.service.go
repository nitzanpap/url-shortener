package urls

import (
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/nitzanpap/url-shortener/server/pkg/utils"
)

func saveUrl(url string, db *pgx.Conn) (string, error) {
	obfuscatedShortenedUrl := shortenAndObfuscateStringUniquely(url, db)

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

func shortenAndObfuscateStringUniquely(url string, db *pgx.Conn) string {
	base62String := utils.GenerateTruncatedHashInBase62(url)

	// Check for collision (pseudo-code)
	for i := 1; checkCollision(base62String, db, url); i++ {
		// If there is a collision, add numbers incrementally to the URL
		url += strconv.Itoa(i)
		base62String = utils.GenerateTruncatedHashInBase62(url)
	}

	return base62String
}

func checkCollision(base62String string, db *pgx.Conn, url string) bool {
	urlFromDb, err := getUrlFromDb(base62String, db)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "no rows in result set" {
			println("No rows in result set")
		}
		return false
	}
	return urlFromDb != url
}
