package service

import (
	"fmt"
	"strings"

	"github.com/antirecord/url-shortener/internal/app/config"
	"github.com/antirecord/url-shortener/internal/app/entity"
)

type SimpleURLShortener struct {
	Config  config.Config
	storage map[string]entity.StorageEntity
}

func NewURLShortener(storage map[string]entity.StorageEntity, config config.Config) *SimpleURLShortener {
	return &SimpleURLShortener{storage: storage, Config: config}
}

func (us SimpleURLShortener) Shorten(url string) (string, error) {
	fmt.Println("url = ", url)
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "", fmt.Errorf("url должен начинаться на http:// или https://")
	}

	hash := GenerateHash(url)
	fmt.Println("hash === ", hash)
	newURL := mergeHash(hash, us.Config.BaseURL)
	fmt.Println("newUrl ==== ", newURL)
	entity := entity.StorageEntity{
		BaseURL:    url,
		ShortedURL: newURL,
	}

	us.storage[hash] = entity
	return newURL, nil
}

func (us SimpleURLShortener) GetBaseURL(id string) (string, error) {
	baseURL, ok := us.storage[id]

	if ok {
		return baseURL.BaseURL, nil
	}
	return "", fmt.Errorf("url с таким id не найден")
}

func mergeHash(hash, baseURL string) string {
	return fmt.Sprintf("http://%s/%s", baseURL, hash)
}
