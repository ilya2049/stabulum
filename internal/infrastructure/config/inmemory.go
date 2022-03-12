package config

import (
	"time"
)

func ReadFromMemory() Config {
	return Config{
		API: APIConfig{
			HTTPServer: APIHTTPServerConfig{
				Address: ":8080",
			},
		},
		ProductUsecases: ProductUsecasesConfig{
			Retry: ProductUsecasesRetryConfig{
				MaxAttempt: 10,
				RetryDelay: time.Second,
			},
		},
		Postgres: PostgresConfig{
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "store-app",
			Password: "password",
			Database: "store-app-db",
		},
	}
}
