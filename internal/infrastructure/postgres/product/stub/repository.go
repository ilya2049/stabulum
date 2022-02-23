package stub

import (
	"context"
	"errors"
	"log"
	"stabulum/internal/domain/product"
)

type ProductRepositoryConfig struct {
	MaxFailedAttempt int
}

type ProductRepository struct {
	attempt int

	cfg ProductRepositoryConfig
}

func NewProductRepository(cfg ProductRepositoryConfig) *ProductRepository {
	r := &ProductRepository{
		cfg: cfg,
	}

	r.resetAttempt()

	return r
}

var errAddProduct = errors.New("failed to add a product to the stub repository")

func (r *ProductRepository) Add(_ context.Context, p product.Product) error {
	if r.attempt > r.cfg.MaxFailedAttempt {
		r.resetAttempt()

		log.Println("product added in the stub repository:", p.String())

		return nil
	}

	r.attempt++

	return errAddProduct
}

func (r *ProductRepository) resetAttempt() {
	r.attempt = 1
}
