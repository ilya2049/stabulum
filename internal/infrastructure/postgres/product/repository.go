package product

import (
	"context"
	"stabulum/internal/common/logger"
	"stabulum/internal/domain/product"
	"stabulum/internal/infrastructure/postgres"
)

type Repository struct {
	logger             logger.Logger
	postgresConnection *postgres.Connection
}

func NewRepository(
	logger logger.Logger,
	postgresConnection *postgres.Connection,
) *Repository {
	return &Repository{
		logger:             logger,
		postgresConnection: postgresConnection,
	}
}

func (r *Repository) Add(_ context.Context, p product.Product) error {
	r.logger.Println("product added in the postgres repository:", p.String())

	return nil
}
