package product

import (
	"context"
	"stabulum/internal/common/event"
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
	return u.productRepository.Atomic(ctx,
		func(ctx context.Context, productRepository product.Repository, eventPublisher event.Publisher) error {
			if err := u.productRepository.Add(ctx, p); err != nil {
				return err
			}

			if err := eventPublisher.Publish(product.CreatedEvent{Name: p.Name}); err != nil {
				return err
			}

			return nil
		})
}
