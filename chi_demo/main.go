package main

import (
	"fmt"
	"log"

	"github.com/moroz/chi_demo/api"
	"github.com/moroz/chi_demo/db"
)

func setUpServices() error {
	_, err := db.ConnectToDb()
	if err != nil {
		return err
	}

	err = api.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := setUpServices()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Hello, world!")
}
