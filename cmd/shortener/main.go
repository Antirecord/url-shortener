package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/antirecord/url-shortener/internal/app/entity"
	"github.com/antirecord/url-shortener/internal/app/service"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	if err := run(mux); err != nil {
		panic(err)
	}
}

func run(mux *http.ServeMux) error {
	return http.ListenAndServe(":8080", mux)
}

var urlShortner = service.NewURLShortener{Storage: make(map[string]entity.StorageEntity)}

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		id := r.URL.Path[1:]
		res, err := urlShortner.GetBaseURL(id)
		if err != nil {
			http.Error(w, "Error while getting baseUrl, url not found", http.StatusNotFound)
		}
		w.Header().Add("Location", res)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error while reading body: %s", err.Error()), http.StatusBadRequest)
			return
		}

		res, err := urlShortner.Shorten(string(body))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error while shorting body: %s", err.Error()), http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(res))
	default:
		http.Error(w, "Specified status not allowed", http.StatusMethodNotAllowed)
	}

}
