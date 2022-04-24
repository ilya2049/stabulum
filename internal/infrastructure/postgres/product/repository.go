package product

import (
	"context"
	"fmt"
	"stabulum/internal/common/event"
	"stabulum/internal/domain/product"
	"stabulum/internal/infrastructure/postgres"
)

type Repository struct {
	db             postgres.DB
	tx             postgres.Tx
	eventPublisher event.Publisher
}

func NewRepository(
	db postgres.DB,
	tx postgres.Tx,
	eventPublisher event.Publisher,
) *Repository {
	return &Repository{
		db:             db,
		tx:             tx,
		eventPublisher: eventPublisher,
	}
}

func (r *Repository) Atomic(ctx context.Context,
	f func(context.Context, product.Repository, event.Publisher) error,
) error {
	tx, err := r.tx.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to open transaction in product repository: %w", err)
	}

	if err := f(ctx, r, r.eventPublisher); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to close transaction in product repository: %w", err)
	}

	return nil
}

func (r *Repository) Add(ctx context.Context, p product.Product) error {
	const query = `INSERT INTO products(name) VALUES ($1);`

	_, err := r.db.ExecContext(ctx, query, p.Name)
	if err != nil {
		return fmt.Errorf("%s: failed to add a product: %w", postgres.Component, err)
	}

	return nil
}
