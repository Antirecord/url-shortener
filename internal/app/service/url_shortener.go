package service

import (
	"fmt"
	"strings"

	"github.com/antirecord/url-shortener/internal/app/entity"
)

type NewUrlShortener struct {
	Storage map[string]entity.StorageEntity
}

func (us NewUrlShortener) Shorten(url string) (string, error) {
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

	us.Storage[hash] = entity
	return newUrl, nil
}

func (us NewUrlShortener) GetBaseUrl(id string) (string, error) {
	baseUrl, ok := us.Storage[id]

	if ok {
		return baseUrl.BaseUrl, nil
	}
	return "", fmt.Errorf("url с таким id не найден")
}

func mergeHash(hash string) string {
	return fmt.Sprintf("http://localhost:8080/%s", hash)
}
