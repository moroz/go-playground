package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

const DEFAULT_PORT string = ":4000"

func getPort() string {
	port := os.Getenv("PORT")
	portNumber, err := strconv.Atoi(port)
	if err != nil {
		return DEFAULT_PORT
	}
	return fmt.Sprintf(":%d", portNumber)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	port := getPort()
	log.Printf("Listening on port %s\n", port)
	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
