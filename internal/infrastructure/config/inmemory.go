package config

import (
	"context"
	"log"
	"time"

	"stabulum/internal/domain/product"
	mockproduct "stabulum/internal/domain/product/mocks"
	"stabulum/internal/testfixture"

	"github.com/stretchr/testify/mock"
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

func ReadFromMemoryMockConfig() MockConfig {
	return MockConfig{
		ConfigureProductRepository: func(r *mockproduct.Repository) {
			const maxFailedAttempt = 5
			attempt := 1

			r.On("Add", mock.Anything, mock.Anything).
				Return(func(_ context.Context, p product.Product) error {
					if attempt >= maxFailedAttempt {
						log.Println("product added in the mock repository:", p.String())
						attempt = 0

						return nil
					}

					attempt++

					return testfixture.ErrTestUnexpected
				})
		},
	}
}
