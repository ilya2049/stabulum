package config

import (
	"stabulum/internal/app/product"
	"stabulum/internal/infrastructure/httpserver"
)

func NewUsecasesConfig(cfg Config) product.UsecasesConfig {
	return product.UsecasesConfig{
		Retry: product.UsecasesRetryConfig{
			MaxAttemtp: cfg.ProductUsecases.Retry.MaxAttemtp,
			RetryDelay: cfg.ProductUsecases.Retry.RetryDelay,
		},
	}
}

func NewHTTPServerConfig(cfg Config) httpserver.Config {
	return httpserver.Config{
		Address: cfg.API.HTTPServer.Address,
	}
}
