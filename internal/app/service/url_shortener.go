package service

import (
	"fmt"
	"strings"

	"github.com/antirecord/url-shortener/internal/app/entity"
)

type SimpleUrlShortener struct {
	storage map[string]entity.StorageEntity
}

func NewUrlShortener(storage map[string]entity.StorageEntity) *SimpleUrlShortener {
	return &SimpleUrlShortener{storage: storage}
}

func (us SimpleUrlShortener) Shorten(url string) (string, error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "", fmt.Errorf("url должен начинаться на http:// или https://")
	}

	hash := GenerateHash(url)
	fmt.Println("hash === ", hash)
	newUrl := mergeHash(hash)
	fmt.Println("newUrl ==== ", newUrl)
	entity := entity.StorageEntity{
		BaseUrl:    url,
		ShortedUrl: newUrl,
	}

	us.storage[hash] = entity
	return newUrl, nil
}

func (us SimpleUrlShortener) GetBaseUrl(id string) (string, error) {
	baseUrl, ok := us.storage[id]

	if ok {
		return baseUrl.BaseUrl, nil
	}
	return "", fmt.Errorf("url с таким id не найден")
}

func mergeHash(hash string) string {
	return fmt.Sprintf("http://localhost:8080/%s", hash)
}
