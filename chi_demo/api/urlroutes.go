package api

import (
	"encoding/json"
	"log"
	"net/http"

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

type shortenURLParams struct {
	Url string `json:"url"`
}

type shortenURLErrorResponse struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

type shortenURLSuccessResponse struct {
	Success bool    `json:"success"`
	Data    *db.URL `json:"data"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var params shortenURLParams

	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, err := db.ShortenURL(params.Url)
	if err != nil {
		log.Println(err)
		response := &shortenURLErrorResponse{
			Success: false,
			Errors:  []string{"validation error"},
		}
		body, _ := json.Marshal(&response)
		http.Error(w, string(body), http.StatusUnprocessableEntity)
		return
	}

	response := &shortenURLSuccessResponse{
		Success: true,
		Data:    record,
	}
	body, _ := json.Marshal(&response)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	w.Write(body)
}
