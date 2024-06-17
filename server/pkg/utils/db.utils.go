package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nitzanpap/url-shortener/server/pkg/colors"
)

func BuildPostgresqlDbURL(host, port, user, pass, name string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, name)
}

func TestDbConnection(db *pgx.Conn) {
	// Test the connection to the database and print the response
	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf(colors.Error("could not ping database: %s\n"), err)
	}
	log.Print(colors.Success("Successfully connected to database\n"))
}

func TestDbPoolConnection(db *pgxpool.Pool) {
	// Test the connection to the database and print the response
	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf(colors.Error("could not ping database: %s\n"), err)
	}
	log.Print(colors.Success("Successfully connected to database\n"))
}
