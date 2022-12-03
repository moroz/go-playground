package main

import (
	"fmt"
	"log"

	"github.com/moroz/chi_demo/api"
	"github.com/moroz/chi_demo/db"
)

func main() {
	connStr := db.GetDBConnectionString()
	_, err := db.ConnectToDb(connStr)
	if err != nil {
		log.Fatalln(err)
	}

	err = api.ListenAndServe(":4000")
	fmt.Println("Hello, world!")
}
