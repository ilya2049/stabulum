package main

import (
	"log"
	appproduct "stabulum/internal/app/product"
	"stabulum/internal/domain/product"
	pgproduct "stabulum/internal/infrastructure/postgres/product"
)

func main() {
	productRepostiory := pgproduct.NewRepostiory()
	productUsecases := appproduct.NewUsecases(productRepostiory)

	if err := productUsecases.Create(product.New("Sticker")); err != nil {
		log.Println("failed to create a product:", err.Error())

		return
	}

	log.Println("done!")
}
