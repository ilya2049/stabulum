package config

import (
	"stabulum/internal/app/product"
)

func NewUsecasesConfig(cfg Config) product.UsecasesConfig {
	return product.UsecasesConfig{
		Retry: product.UsecasesRetryConfig{
			MaxAttemtp: cfg.ProductUsecasesConfig.Retry.MaxAttemtp,
			RetryDelay: cfg.ProductUsecasesConfig.Retry.RetryDelay,
		},
	}
}
