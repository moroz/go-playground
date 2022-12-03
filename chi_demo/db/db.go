package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const DEFAULT_CONN_STRING = "postgres://postgres:postgres@localhost:5432/url_shortener_dev?sslmode=disable"

func GetDBConnectionString() string {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return DEFAULT_CONN_STRING
	}
	return connStr
}

func ConnectToDb(connStr string) (*sql.DB, error) {
	log.Println("Connecting to PostgreSQL database...")
	db, err := sql.Open("postgres", connStr)
	return db, err
}
