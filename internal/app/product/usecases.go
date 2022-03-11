package product

import (
	"context"
	"stabulum/internal/domain/product"
)

type usecases struct {
	productRepository product.Repository
}

func (u *usecases) Create(ctx context.Context, p product.Product) error {
	return u.productRepository.Add(ctx, p)
}
