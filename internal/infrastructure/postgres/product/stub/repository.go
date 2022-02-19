package stub

import (
	"context"
	"errors"
	"log"
	"stabulum/internal/domain/product"
)

type ProductRepository struct {
	attempt int
}

func NewProductRepository() *ProductRepository {
	r := &ProductRepository{}
	r.resetAttempt()

	return r
}

const maxFailedAttempt = 14

var errAddProduct = errors.New("failed to add a product to the stub repository")

func (r *ProductRepository) Add(_ context.Context, p product.Product) error {
	if r.attempt > maxFailedAttempt {
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
