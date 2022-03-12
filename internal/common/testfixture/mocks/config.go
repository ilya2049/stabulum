package mocks

import (
	"stabulum/internal/common/logger"
	mockproduct "stabulum/internal/domain/product/mocks"
)

type Config struct {
	ConfigureProductRepository func(*mockproduct.Repository, logger.Logger)
}
