package product

import (
	"context"
	"database/sql"
	"stabulum/internal/common/logger"
	"stabulum/internal/domain/product"
)

type Repository struct {
	logger             logger.Logger
	postgresConnection *sql.DB
}

func NewRepository(
	logger logger.Logger,
	postgresConnection *sql.DB,
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
