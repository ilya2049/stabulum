//go:build wireinject
// +build wireinject

package di

import (
	appproduct "stabulum/internal/app/product"
	"stabulum/internal/domain/product"
	pgproduct "stabulum/internal/infrastructure/postgres/product"
	productstub "stabulum/internal/infrastructure/postgres/product/stub"

	"github.com/google/wire"
)

func NewContainer() *Container {
	panic(
		wire.Build(
			newContainer,

			appproduct.NewUsecases,

			productPostgresRepositorySet,
		),
	)
}

var productPostgresRepositorySet = wire.NewSet(
	wire.Bind(new(product.Repository), new(*pgproduct.Repository)),
	pgproduct.NewRepostiory,
)

func NewTestContainer() *Container {
	panic(
		wire.Build(
			newContainer,

			appproduct.NewUsecases,

			productStubRepositorySet,
		),
	)
}

var productStubRepositorySet = wire.NewSet(
	wire.Bind(new(product.Repository), new(*productstub.ProductRepository)),
	productstub.NewProductRepository,
)
