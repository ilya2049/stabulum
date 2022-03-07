//go:build wireinject
// +build wireinject

package di

import (
	"net/http"

	appproduct "stabulum/internal/app/product"
	"stabulum/internal/domain/product"
	mockproduct "stabulum/internal/domain/product/mocks"
	"stabulum/internal/infrastructure/api/router"
	apiproduct "stabulum/internal/infrastructure/api/router/product"
	"stabulum/internal/infrastructure/config"
	"stabulum/internal/infrastructure/httpserver"
	pgproduct "stabulum/internal/infrastructure/postgres/product"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func NewContainer(cfg config.Config) *Container {
	panic(
		wire.Build(
			newContainer,
			configSet,

			apiSet,

			appproduct.NewUsecases,

			productPostgresRepositorySet,
		),
	)
}

var configSet = wire.NewSet(
	config.NewUsecasesConfig,
	config.NewHTTPServerConfig,
)

var apiSet = wire.NewSet(
	wire.Bind(new(http.Handler), new(*gin.Engine)),
	httpserver.New,
	router.New,

	apiproduct.NewHandler,
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

			apiSet,

			appproduct.NewUsecases,

			productMockRepositorySet,
		),
	)
}

var productMockRepositorySet = wire.NewSet(
	wire.Bind(new(product.Repository), new(*mockproduct.Repository)),
	config.NewProductRepositoryMock,
)
