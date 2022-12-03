package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moroz/chi_demo/db"
)

const DEFAULT_PORT = 4000

func ListURLs(w http.ResponseWriter, r *http.Request) {
	urls := db.ListURLs()
	result, err := json.Marshal(urls)
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	w.Write(result)
}

func BuildHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/urls", ListURLs)

	return r
}

func ResolvePort() int {
	portEnv := os.Getenv("PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		return DEFAULT_PORT
	}
	return port
}

func ListenAndServe() error {
	port := ResolvePort()
	addrString := fmt.Sprintf(":%d", port)

	server := http.Server{
		Addr:    addrString,
		Handler: BuildHandler(),
	}

	log.Printf("Listening on port %d\n", port)

	err := server.ListenAndServe()
	return err
}
