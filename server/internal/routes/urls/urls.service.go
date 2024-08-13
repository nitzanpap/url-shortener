package urls

import (
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/nitzanpap/url-shortener/server/pkg"
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
	// Check if the original URL already exists in the database
	existingBase62String, err := getBase62StringFromDb(url, db)
	if err == nil && existingBase62String != "" {
		// If the URL already exists in the DB, return the associated base62String
		return existingBase62String
	}

	// If the URL is new, proceed with generating a new base62String
	base62String := utils.GenerateTruncatedHashInBase62(url, pkg.NUM_OF_CHARS_IN_URL_ID)

	// Check for collision
	originalUrl := url
	for i := 1; checkCollision(base62String, db, originalUrl) && i < pkg.NumOfPossibleUrls; i++ {
		// Modify the URL with an increment
		url = originalUrl + strconv.Itoa(i)
		base62String = utils.GenerateTruncatedHashInBase62(url, pkg.NUM_OF_CHARS_IN_URL_ID)
	}

	return base62String
}

func checkCollision(base62String string, db *pgx.Conn, url string) bool {
	urlFromDb, err := getUrlFromDb(base62String, db)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "no rows in result set" {
			println("No rows in result set")
			return false
		}
	}
	return urlFromDb != "" && urlFromDb != url
}
