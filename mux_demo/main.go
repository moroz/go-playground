package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const DEFAULT_PORT = ":4000"

func getServerAddr() string {
	portVar := os.Getenv("PORT")
	port, err := strconv.Atoi(portVar)
	if err != nil {
		return DEFAULT_PORT
	}
	return fmt.Sprintf(":%d", port)
}

func buildServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    getServerAddr(),
		Handler: handler,
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"status":"success"}`)
}

func buildRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", helloWorldHandler)

	handler := handlers.LoggingHandler(os.Stdout, r)
	return handler
}

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	router := buildRouter()
	server := buildServer(router)

	log.Printf("Listening on %s\n", server.Addr)
	server.ListenAndServe()
}
