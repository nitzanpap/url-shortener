package urls

import (
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/nitzanpap/url-shortener/server/pkg"
	"github.com/nitzanpap/url-shortener/server/pkg/utils"
)

func saveURL(url string, db *pgx.Conn) (string, error) {
	obfuscatedShortenedURL := shortenAndObfuscateStringUniquely(url, db)

	// save the URL and the obfuscatedShortenedURL in the database
	err := saveURLInDB(url, obfuscatedShortenedURL, db)
	if err != nil {
		return "", err
	}
	return obfuscatedShortenedURL, nil
}

func getURL(obfuscatedShortenedURL string, db *pgx.Conn) (string, error) {
	// get the URL from the database using the obfuscatedShortenedURL
	url, err := getURLFromDB(obfuscatedShortenedURL, db)
	if err != nil {
		return "", err
	}
	return url, nil
}

func shortenAndObfuscateStringUniquely(url string, db *pgx.Conn) string {
	// Check if the original URL already exists in the database
	existingBase62String, err := getBase62StringFromDB(url, db)
	if err == nil && existingBase62String != "" {
		// If the URL already exists in the DB, return the associated base62String
		return existingBase62String
	}

	// If the URL is new, proceed with generating a new base62String
	base62String := utils.GenerateTruncatedHashInBase62(url, pkg.NumOfCharsInURLID)

	// Check for collision
	originalURL := url
	for i := 1; checkCollision(base62String, db, originalURL) && i < pkg.NumOfPossibleUrls; i++ {
		// Modify the URL with an increment
		url = originalURL + strconv.Itoa(i)
		base62String = utils.GenerateTruncatedHashInBase62(url, pkg.NumOfCharsInURLID)
	}

	return base62String
}

func checkCollision(base62String string, db *pgx.Conn, url string) bool {
	urlFromDB, err := getURLFromDB(base62String, db)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "no rows in result set" {
			println("No rows in result set")
			return false
		}
	}
	return urlFromDB != "" && urlFromDB != url
}
