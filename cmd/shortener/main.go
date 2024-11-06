package main

import (
	"net/http"

	"github.com/antirecord/url-shortener/internal/app/entity"
	"github.com/antirecord/url-shortener/internal/app/service"
	"github.com/antirecord/url-shortener/internal/handlers"
)

func main() {
	urlShortener := service.NewUrlShortener(map[string]entity.StorageEntity{})
	handler := handlers.NewHandler(urlShortener)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Handle)

	if err := run(mux); err != nil {
		panic(err)
	}
}

func run(mux *http.ServeMux) error {
	return http.ListenAndServe(":8080", mux)
}
