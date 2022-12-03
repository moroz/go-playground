package server

import (
	"fmt"
	"log"
	"moroz/chi_demo/db"
)

func main() {
	connStr := db.GetDBConnectionString()
	db, err := db.ConnectToDb(connStr)
	_ = db
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Hello, world!")
}
