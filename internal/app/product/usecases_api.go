package product

import (
	"context"
	"stabulum/internal/domain/product"
	"time"
)

type Usecases interface {
	Create(context.Context, product.Product) error
}

func NewUsecases(config UsecasesConfig, productRepostiory product.Repository) Usecases {
	var uc Usecases

	uc = &usecases{
		productRepostiory: productRepostiory,
	}

	uc = &usecasesRetry{
		next: uc,

		config: config.Retry,
	}

	return uc
}

type UsecasesConfig struct {
	Retry UsecasesRetryConfig
}

type UsecasesRetryConfig struct {
	MaxAttemtp int
	RetryDelay time.Duration
}
