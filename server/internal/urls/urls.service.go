package urls

import (
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/nitzanpap/url-shortener/server/pkg/colors"
	"github.com/nitzanpap/url-shortener/server/pkg/utils"
)

func saveUrl(url string, db *pgx.Conn) (string, error) {
	// generate a obfuscatedShortenedUrl for the URL
	obfuscatedShortenedUrl := shortenAndObfuscateStringUniquely(url)

	// save the URL and the obfuscatedShortenedUrl in the database
	err := saveUrlInDb(url, obfuscatedShortenedUrl, db)
	if err != nil {
		return "", err
	}

	// return the shortened URL
	return obfuscatedShortenedUrl, nil
}

func getUrl(obfuscatedShortenedUrl string, db *pgx.Conn) (string, error) {
	// get the URL from the database using the obfuscatedShortenedUrl
	url, err := getUrlFromDb(obfuscatedShortenedUrl, db)
	if err != nil {
		return "", err
	}

	// return the URL
	return url, nil
}

func shortenAndObfuscateStringUniquely(url string) string {
	base62String := utils.Base62Encode([]byte(url))
	log.Printf(colors.Info("Original URL: %s"), url)
	log.Printf(colors.Info("Shortened URL: %s"), base62String)
	return base62String
}
