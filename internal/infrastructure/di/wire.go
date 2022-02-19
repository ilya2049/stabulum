//go:build wireinject
// +build wireinject

package di

import (
	appproduct "stabulum/internal/app/product"
	"stabulum/internal/domain/product"
	pgproduct "stabulum/internal/infrastructure/postgres/product"

	"github.com/google/wire"
)

func NewContainer() *Container {
	panic(
		wire.Build(
			newContainer,

			appproduct.NewUsecases,

			wire.Bind(new(product.Repository), new(*pgproduct.Repository)),
			pgproduct.NewRepostiory,
		),
	)
}
