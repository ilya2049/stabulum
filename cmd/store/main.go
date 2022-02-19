package main

import (
	"context"
	"log"
	"stabulum/internal/domain/product"
	"stabulum/internal/infrastructure/di"
)

func main() {
	diContainer := di.NewTestContainer()

	productUsecases := diContainer.ProductUsecases

	if err := productUsecases.Create(context.Background(), product.New("Sticker")); err != nil {
		log.Println("failed to create a product:", err.Error())

		return
	}

	log.Println("done!")
}
