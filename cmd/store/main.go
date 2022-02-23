package main

import (
	"context"
	"log"
	"stabulum/internal/domain/product"
	"stabulum/internal/infrastructure/config"
	"stabulum/internal/infrastructure/di"
	"time"
)

func main() {
	diContainer := di.NewTestContainer(config.Config{
		ProductUsecasesConfig: config.ProductUsecasesConfig{
			Retry: config.ProductUsecasesRetryConfig{
				MaxAttemtp: 10,
				RetryDelay: time.Second,
			},
		},
		ProductRepositoryStubConfig: config.ProductRepositoryStubConfig{
			MaxFailedAttempt: 14,
		},
	})

	productUsecases := diContainer.ProductUsecases

	if err := productUsecases.Create(context.Background(), product.New("Sticker")); err != nil {
		log.Println("failed to create a product:", err.Error())

		return
	}

	log.Println("done!")
}
