package config

import "github.com/maximprokopchuk/auto_reference_catalog_service/internal/store"

type Config struct {
	BindUrl string `toml:"bind_url"`
	Store   *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindUrl: "8080",
		Store:   store.NewConfig(),
	}
}
