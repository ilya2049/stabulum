package mocks

import (
	mockqueries "stabulum/internal/app/queries/mocks"
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

func NewProductQuerierMock(cfg Config, logger logger.Logger) *mockqueries.ProductQuerier {
	q := mockqueries.ProductQuerier{}

	if cfg.ConfigureProductQuerier != nil {
		cfg.ConfigureProductQuerier(&q, logger)
	}

	return &q
}
