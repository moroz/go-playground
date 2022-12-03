package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moroz/chi_demo/db"
)

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

func ListenAndServe(port string) error {
	server := http.Server{
		Addr:    port,
		Handler: BuildHandler(),
	}

	log.Printf("Listening on port %s\n", port)

	err := server.ListenAndServe()
	return err
}
