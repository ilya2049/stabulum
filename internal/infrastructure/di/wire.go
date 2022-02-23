//go:build wireinject
// +build wireinject

package di

import (
	appproduct "stabulum/internal/app/product"
	"stabulum/internal/domain/product"
	"stabulum/internal/infrastructure/config"
	pgproduct "stabulum/internal/infrastructure/postgres/product"
	productstub "stabulum/internal/infrastructure/postgres/product/stub"

	"github.com/google/wire"
)

func NewContainer(cfg config.Config) *Container {
	panic(
		wire.Build(
			newContainer,
			configSet,

			appproduct.NewUsecases,

			productPostgresRepositorySet,
		),
	)
}

var configSet = wire.NewSet(
	config.NewUsecasesConfig,
	config.NewProductRepositoryStubConfig,
)

var productPostgresRepositorySet = wire.NewSet(
	wire.Bind(new(product.Repository), new(*pgproduct.Repository)),
	pgproduct.NewRepostiory,
)

func NewTestContainer(cfg config.Config) *Container {
	panic(
		wire.Build(
			newContainer,
			configSet,

			appproduct.NewUsecases,

			productStubRepositorySet,
		),
	)
}

var productStubRepositorySet = wire.NewSet(
	wire.Bind(new(product.Repository), new(*productstub.ProductRepository)),
	productstub.NewProductRepository,
)
