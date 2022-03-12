package config

import (
	"stabulum/internal/app/product"
	"stabulum/internal/infrastructure/httpserver"
	"stabulum/internal/infrastructure/postgres"
)

func NewUsecasesConfig(cfg Config) product.UsecasesConfig {
	return product.UsecasesConfig{
		Retry: product.UsecasesRetryConfig{
			MaxAttempt: cfg.ProductUsecases.Retry.MaxAttempt,
			RetryDelay: cfg.ProductUsecases.Retry.RetryDelay,
		},
	}
}

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
