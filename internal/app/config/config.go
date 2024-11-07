package config

import "flag"

type Config struct {
	BaseURL string
	Addr    string
}

func NewConfig() *Config {
	var baseURL, addr string
	flag.StringVar(&baseURL, "b", "localhost:8000", "base url for url shortened")
	flag.StringVar(&addr, "a", "localhost:8888", "addr for running server")
	flag.Parse()
	return &Config{BaseURL: baseURL, Addr: addr}
}
