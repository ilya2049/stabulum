package product

import (
	"context"
	"stabulum/internal/domain/product"
)

type Usecases interface {
	Create(context.Context, product.Product) error
}

func NewUsecases(productRepository product.Repository) Usecases {
	return &usecases{
		productRepository: productRepository,
	}
}

type usecases struct {
	productRepository product.Repository
}

func (u *usecases) Create(ctx context.Context, p product.Product) error {
	return u.productRepository.Add(ctx, p)
}
