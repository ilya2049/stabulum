package config

import (
	"stabulum/internal/app/product"
	"stabulum/internal/infrastructure/postgres/product/stub"
)

func NewUsecasesConfig(cfg Config) product.UsecasesConfig {
	return product.UsecasesConfig{
		Retry: product.UsecasesRetryConfig{
			MaxAttemtp: cfg.ProductUsecasesConfig.Retry.MaxAttemtp,
			RetryDelay: cfg.ProductUsecasesConfig.Retry.RetryDelay,
		},
	}
}

func NewProductRepositoryStubConfig(cfg Config) stub.ProductRepositoryConfig {
	return stub.ProductRepositoryConfig{
		MaxFailedAttempt: cfg.ProductRepositoryStubConfig.MaxFailedAttempt,
	}
}
