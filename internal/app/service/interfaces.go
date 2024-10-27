package service

type UrlShortener interface {
	Shorten(data string) (string, error)
}
