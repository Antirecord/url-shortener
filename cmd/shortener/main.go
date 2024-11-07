package main

import (
	"net/http"

	"github.com/antirecord/url-shortener/internal/app/config"
	"github.com/antirecord/url-shortener/internal/app/entity"
	"github.com/antirecord/url-shortener/internal/app/handlers"
	"github.com/antirecord/url-shortener/internal/app/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	conf := *config.NewConfig()
	urlShortener := service.NewURLShortener(map[string]entity.StorageEntity{}, conf)

	handler := handlers.NewHandler(urlShortener)

	r := chi.NewRouter()
	r.Post("/", handler.Handle)
	r.Get("/{id}", handler.Handle)
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", handler.Handle)

	if err := run(r, conf.Addr); err != nil {
		panic(err)
	}
}

func run(mux http.Handler, addr string) error {
	return http.ListenAndServe(addr, mux)
}
