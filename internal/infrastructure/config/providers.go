package config

import (
	"stabulum/internal/infrastructure/httpserver"
	"stabulum/internal/infrastructure/postgres"
)

func NewHTTPServerConfig(cfg Config) httpserver.Config {
	return httpserver.Config{
		Address: cfg.API.HTTPServer.Address,
	}
}

func NewPostgresConfig(cfg Config) postgres.Config {
	return postgres.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		Database: cfg.Postgres.Database,
	}
}
