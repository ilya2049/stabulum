package product

import (
	"context"
	"log"
	"stabulum/internal/domain/product"
	"stabulum/internal/infrastructure/postgres"
)

type Repository struct {
	postgresConnection *postgres.Connection
}

func NewRepository(postgresConnection *postgres.Connection) *Repository {
	return &Repository{
		postgresConnection: postgresConnection,
	}
}

func (r *Repository) Add(_ context.Context, p product.Product) error {
	log.Println("product added in the postgres repository:", p.String())

	return nil
}
