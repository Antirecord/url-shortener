package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/antirecord/url-shortener/internal/app/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockURLShortener struct {
	storage map[string]entity.StorageEntity
}

func (m *mockURLShortener) GetBaseURL(id string) (string, error) {
	url, exists := m.storage[id]

	if !exists {
		return "", fmt.Errorf("url not found")
	}
	return url.BaseURL, nil
}

func (m *mockURLShortener) Shorten(url string) (string, error) {
	id := "short-id"

	m.storage[id] = entity.StorageEntity{
		BaseURL:    url,
		ShortedURL: "http://localhost:8080/" + id,
	}
	return id, nil
}

func newTestHandler() *Handler {
	return NewHandler(&mockURLShortener{map[string]entity.StorageEntity{}})
}

func TestHandlerGet(t *testing.T) {
	handler := newTestHandler()
	url := "https://ya.ru"

	handler.UrlShortener.(*mockURLShortener).storage["short-id"] = entity.StorageEntity{BaseURL: url}
	req := httptest.NewRequest(http.MethodGet, "/short-id", nil)
	w := httptest.NewRecorder()
	handler.Handle(w, req)

	res := w.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Fatalf("Expected status %d, got %d", http.StatusTemporaryRedirect, res.StatusCode)
	}

	location := res.Header.Get("Location")
	if location != url {
		t.Errorf("Expected location %s, got %s", url, location)
	}
}

func TestHandlerGetNotFound(t *testing.T) {
	handler := newTestHandler()

	req := httptest.NewRequest(http.MethodGet, "/fail-id", nil)
	w := httptest.NewRecorder()
	handler.Handle(w, req)

	res := w.Result()

	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected status %d, got %d", http.StatusNotFound, res.StatusCode)
	}
}

func TestHandlerPost(t *testing.T) {
	handler := newTestHandler()
	url := "https://ya.ru"
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(url))
	w := httptest.NewRecorder()
	handler.Handle(w, req)

	res := w.Result()

	defer res.Body.Close()
	assert.Equal(t, res.StatusCode, http.StatusCreated)

	body, err := io.ReadAll(res.Body)

	require.NoError(t, err)
	assert.Equal(t, string(body), "short-id")
}
func TestHandlerPostBadRequest(t *testing.T) {
	handler := newTestHandler()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	handler.Handle(w, req)

	res := w.Result()

	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
}

func TestHandlerMethodNotAllowed(t *testing.T) {
	handler := newTestHandler()

	req := httptest.NewRequest(http.MethodPut, "/", nil)
	w := httptest.NewRecorder()
	handler.Handle(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, http.StatusMethodNotAllowed)
}
