package mocks

import (
	mockproduct "stabulum/internal/domain/product/mocks"
)

type Config struct {
	ConfigureProductRepository func(*mockproduct.Repository)
}
