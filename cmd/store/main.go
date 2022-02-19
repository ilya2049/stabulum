package main

import (
	"log"
	"stabulum/internal/domain/product"
	"stabulum/internal/infrastructure/di"
)

func main() {
	diContainer := di.NewContainer()

	productUsecases := diContainer.ProductUsecases

	if err := productUsecases.Create(product.New("Sticker")); err != nil {
		log.Println("failed to create a product:", err.Error())

		return
	}

	log.Println("done!")
}
