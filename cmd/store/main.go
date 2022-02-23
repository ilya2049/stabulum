package main

import (
	"context"
	"log"
	"stabulum/internal/domain/product"
	mockproduct "stabulum/internal/domain/product/mocks"
	"stabulum/internal/infrastructure/config"
	"stabulum/internal/infrastructure/di"
	"stabulum/internal/testfixture"
	"time"

	"github.com/stretchr/testify/mock"
)

func main() {
	diContainer := di.NewTestContainer(config.Config{
		ProductUsecasesConfig: config.ProductUsecasesConfig{
			Retry: config.ProductUsecasesRetryConfig{
				MaxAttemtp: 10,
				RetryDelay: time.Second,
			},
		},
	},
		config.MockConfig{
			ConfigureProductRepository: func(r *mockproduct.Repository) {
				const maxFailedAttempt = 5
				attempt := 1

				r.On("Add", mock.Anything, mock.Anything).
					Return(func(_ context.Context, p product.Product) error {
						if attempt >= maxFailedAttempt {
							log.Println("product added in the mock repository:", p.String())

							return nil
						}

						attempt++

						return testfixture.ErrTestUnexpected
					})
			},
		},
	)

	productUsecases := diContainer.ProductUsecases

	if err := productUsecases.Create(context.Background(), product.New("Sticker")); err != nil {
		log.Println("failed to create a product:", err.Error())

		return
	}

	log.Println("done!")
}
