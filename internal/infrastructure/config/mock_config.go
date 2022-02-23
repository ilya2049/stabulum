package config

import (
	mockproduct "stabulum/internal/domain/product/mocks"
)

type MockConfig struct {
	ConfigureProductRepository func(*mockproduct.Repository)
}
