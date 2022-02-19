package product

import (
	"log"
	"stabulum/internal/domain/product"
)

type Repository struct {
}

func NewRepostiory() *Repository {
	return &Repository{}
}

func (r *Repository) Add(p product.Product) error {
	log.Println("product added in the postgres repository:", p.String())

	return nil
}
