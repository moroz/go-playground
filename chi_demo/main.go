package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost:5432/url_shortener_dev?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Query("Select now() at time zone 'utc';")
	if err != nil {
		log.Printf("Could not execute query: %s\n", err)
	} else {
		var rows []string
		var row string

		for result.Next() {
			result.Scan(&row)
			fmt.Println(row)
			rows = append(rows, row)
		}

		if result.Err() != nil {
			fmt.Println(result.Err())
		}
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	http.ListenAndServe(":4000", r)
}
