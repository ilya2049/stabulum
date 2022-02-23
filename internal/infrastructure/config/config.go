package config

import (
	"time"
)

type Config struct {
	ProductUsecasesConfig       ProductUsecasesConfig
	ProductRepositoryStubConfig ProductRepositoryStubConfig
}

type ProductUsecasesConfig struct {
	Retry ProductUsecasesRetryConfig
}

type ProductUsecasesRetryConfig struct {
	MaxAttemtp int
	RetryDelay time.Duration
}

type ProductRepositoryStubConfig struct {
	MaxFailedAttempt int
}
