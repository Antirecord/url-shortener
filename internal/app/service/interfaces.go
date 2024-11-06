package service

type UrlShortener interface {
	Shorten(data string) (string, error)
	GetBaseUrl(id string) (string, error)
}
