package product

import (
	"context"
	"stabulum/internal/domain/product"
)

type Usecases interface {
	Create(context.Context, product.Product) error
}

func NewUsecases(productRepostiory product.Repository) Usecases {
	var uc Usecases

	uc = &usecases{
		productRepostiory: productRepostiory,
	}

	uc = &usecasesRetry{
		next: uc,
	}

	return uc
}
