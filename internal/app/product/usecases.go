package product

import (
	"context"
	"stabulum/internal/domain/product"
)

type usecases struct {
	productRepostiory product.Repository
}

func (u *usecases) Create(ctx context.Context, p product.Product) error {
	return u.productRepostiory.Add(ctx, p)
}
