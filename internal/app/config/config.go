package config

import "flag"

type Config struct {
	BaseURL string
	Addr    string
}

func NewConfig() *Config {
	var baseURL, addr string
	flag.StringVar(&addr, "a", "localhost:8080", "addr for running server")
	flag.StringVar(&baseURL, "b", "localhost:8000", "base url for url shortened")
	flag.Parse()
	return &Config{BaseURL: baseURL, Addr: addr}
}
