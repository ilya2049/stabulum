package mocks

import (
	"stabulum/internal/common/logger"
	mockproduct "stabulum/internal/domain/product/mocks"
)

func NewProductRepositoryMock(cfg Config, logger logger.Logger) *mockproduct.Repository {
	r := mockproduct.Repository{}

	if cfg.ConfigureProductRepository != nil {
		cfg.ConfigureProductRepository(&r, logger)
	}

	return &r
}
