package service

type URLShortener interface {
	Shorten(data string) (string, error)
	GetBaseURL(id string) (string, error)
}
