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
	}
}
