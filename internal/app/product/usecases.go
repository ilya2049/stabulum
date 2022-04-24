package product

import (
	"context"
	"stabulum/internal/common/event"
	"stabulum/internal/domain/product"
)

type Usecases interface {
	Create(context.Context, product.Product) error
}

func NewUsecases(
	productRepository product.Repository,
	eventPublisher event.Publisher,
) Usecases {
	return &usecases{
		productRepository: productRepository,
		eventPublisher:    eventPublisher,
	}
}

type usecases struct {
	productRepository product.Repository

	eventPublisher event.Publisher
}

func (u *usecases) Create(ctx context.Context, p product.Product) error {
	if err := u.productRepository.Add(ctx, p); err != nil {
		return err
	}

	_ = u.eventPublisher.Publish(product.CreatedEvent{Name: p.Name})

	return nil
}
