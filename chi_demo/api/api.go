package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func BuildHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

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
