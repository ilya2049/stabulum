package config

import (
	"time"
)

type Config struct {
	ProductUsecasesConfig ProductUsecasesConfig
}

type ProductUsecasesConfig struct {
	Retry ProductUsecasesRetryConfig
}

type ProductUsecasesRetryConfig struct {
	MaxAttemtp int
	RetryDelay time.Duration
}
