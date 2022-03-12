package config

import (
	"time"
)

type Config struct {
	API             APIConfig
	ProductUsecases ProductUsecasesConfig
	Postgres        PostgresConfig
}

type ProductUsecasesConfig struct {
	Retry ProductUsecasesRetryConfig
}

type ProductUsecasesRetryConfig struct {
	MaxAttempt int
	RetryDelay time.Duration
}

type APIConfig struct {
	HTTPServer APIHTTPServerConfig
}

type APIHTTPServerConfig struct {
	Address string
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
