package product

import (
	"context"
	"database/sql"
	"fmt"
	"stabulum/internal/domain/product"
	"stabulum/internal/infrastructure/postgres"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
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
