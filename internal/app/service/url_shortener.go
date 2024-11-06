package service

import (
	"fmt"
	"strings"

	"github.com/antirecord/url-shortener/internal/app/entity"
)

type NewURLShortener struct {
	Storage map[string]entity.StorageEntity
}

func (us NewURLShortener) Shorten(URL string) (string, error) {
	if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
		return "", fmt.Errorf("url должен начинаться на http:// или https://")
	}

	hash := GenerateHash(URL)
	fmt.Println("hash === ", hash)
	newURL := mergeHash(hash)
	fmt.Println("newUrl ==== ", newURL)
	entity := entity.StorageEntity{
		BaseURL:    URL,
		ShortedURL: newURL,
	}

	us.Storage[hash] = entity
	return newURL, nil
}

func (us NewURLShortener) GetBaseURL(id string) (string, error) {
	baseURL, ok := us.Storage[id]

	if ok {
		return baseURL.BaseURL, nil
	}
	return "", fmt.Errorf("url с таким id не найден")
}

func mergeHash(hash string) string {
	return fmt.Sprintf("http://localhost:8080/%s", hash)
}
