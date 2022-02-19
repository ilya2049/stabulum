package product

import (
	"context"
	"log"
	"stabulum/internal/domain/product"
)

type Repository struct {
}

func NewRepostiory() *Repository {
	return &Repository{}
}

func (r *Repository) Add(_ context.Context, p product.Product) error {
	log.Println("product added in the postgres repository:", p.String())

	return nil
}
