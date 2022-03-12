package product

import (
	"context"
	"stabulum/internal/common/logger"
	"stabulum/internal/domain/product"
	"time"
)

type Usecases interface {
	Create(context.Context, product.Product) error
}

func NewUsecases(
	config UsecasesConfig,
	logger logger.Logger,
	productRepository product.Repository,
) Usecases {
	var uc Usecases

	uc = &usecases{
		productRepository: productRepository,
	}

	uc = &usecasesRetry{
		next:   uc,
		logger: logger,
		config: config.Retry,
	}

	return uc
}

type UsecasesConfig struct {
	Retry UsecasesRetryConfig
}

type UsecasesRetryConfig struct {
	MaxAttempt int
	RetryDelay time.Duration
}
