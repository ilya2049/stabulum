package mocks

import (
	mockqueries "stabulum/internal/app/queries/mocks"
	"stabulum/internal/common/logger"
	mockproduct "stabulum/internal/domain/product/mocks"
)

type Config struct {
	ConfigureProductRepository func(*mockproduct.Repository, logger.Logger)
	ConfigureProductQuerier    func(*mockqueries.ProductQuerier, logger.Logger)
}
