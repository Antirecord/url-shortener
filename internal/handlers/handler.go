package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/antirecord/url-shortener/internal/app/service"
)

type Handler struct {
	UrlShortener service.UrlShortener
}
type Server struct {
}

func NewHandler(urlshortener service.UrlShortener) *Handler {
	return &Handler{UrlShortener: urlshortener}
}
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		id := r.URL.Path[1:]
		res, err := h.UrlShortener.GetBaseUrl(id)
		if err != nil {
			http.Error(w, "Error while getting baseUrl, url not found", http.StatusNotFound)
			return
		}
		w.Header().Add("Location", res)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case http.MethodPost:
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error while reading body: %s", err.Error()), http.StatusBadRequest)
			return
		}

		if len(body) == 0 {
			http.Error(w, "Body is empty.", http.StatusBadRequest)
			return
		}

		res, err := h.UrlShortener.Shorten(string(body))
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
