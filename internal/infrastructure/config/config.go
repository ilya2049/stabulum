package config

import (
	"time"
)

type Config struct {
	API             APIConfig
	ProductUsecases ProductUsecasesConfig
}

type ProductUsecasesConfig struct {
	Retry ProductUsecasesRetryConfig
}

type ProductUsecasesRetryConfig struct {
	MaxAttemtp int
	RetryDelay time.Duration
}

type APIConfig struct {
	HTTPServer APIHTTPServerConfig
}

type APIHTTPServerConfig struct {
	Address string
}
