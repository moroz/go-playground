package db

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DEFAULT_CONN_STRING = "postgres://postgres:postgres@localhost:5432/url_shortener_dev?sslmode=disable"

var DB *gorm.DB

func getDBConnectionString() string {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatalln("FATAL: Environment variable DATABASE_URL is not set!")
	}
	return connStr
}

func ConnectToDb() (*gorm.DB, error) {
	log.Println("Connecting to PostgreSQL database...")
	connStr := getDBConnectionString()
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	DB = db
	return db, err
}
