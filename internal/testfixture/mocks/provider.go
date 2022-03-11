package mocks

import (
	mockproduct "stabulum/internal/domain/product/mocks"
)

func NewProductRepositoryMock(cfg Config) *mockproduct.Repository {
	r := mockproduct.Repository{}

	if cfg.ConfigureProductRepository != nil {
		cfg.ConfigureProductRepository(&r)
	}

	return &r
}
