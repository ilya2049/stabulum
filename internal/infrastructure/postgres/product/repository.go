package product

import (
	"context"
	"database/sql"
	"fmt"
	"stabulum/internal/common/logger"
	"stabulum/internal/domain/product"
	"stabulum/internal/infrastructure/postgres"
)

type Repository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewRepository(logger logger.Logger, db *sql.DB) *Repository {
	return &Repository{
		logger: logger,
		db:     db,
	}
}

func (r *Repository) Add(ctx context.Context, p product.Product) error {
	const query = `INSERT INTO products(name) VALUES ($1);`

	_, err := r.db.ExecContext(ctx, query, p.Name)
	if err != nil {
		return fmt.Errorf("%s: failed to add a product: %w", postgres.Component, err)
	}

	return nil
}
