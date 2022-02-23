//go:build wireinject
// +build wireinject

package di

import (
	appproduct "stabulum/internal/app/product"
	"stabulum/internal/domain/product"
	mockproduct "stabulum/internal/domain/product/mocks"
	"stabulum/internal/infrastructure/config"
	pgproduct "stabulum/internal/infrastructure/postgres/product"

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
)

var productPostgresRepositorySet = wire.NewSet(
	wire.Bind(new(product.Repository), new(*pgproduct.Repository)),
	pgproduct.NewRepostiory,
)

func NewTestContainer(cfg config.Config, mockCfg config.MockConfig) *Container {
	panic(
		wire.Build(
			newContainer,
			configSet,

			appproduct.NewUsecases,

			productMockRepositorySet,
		),
	)
}

var productMockRepositorySet = wire.NewSet(
	wire.Bind(new(product.Repository), new(*mockproduct.Repository)),
	config.NewProductRepositoryMock,
)
