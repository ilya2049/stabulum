package di

import (
	"stabulum/internal/app/product"
)

type Container struct {
	ProductUsecases product.Usecases
}

func newContainer(productUsecases product.Usecases) *Container {
	return &Container{
		ProductUsecases: productUsecases,
	}
}
