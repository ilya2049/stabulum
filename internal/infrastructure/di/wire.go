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
	"stabulum/internal/infrastructure/postgres"
	pgproduct "stabulum/internal/infrastructure/postgres/product"
	"stabulum/internal/pkg/connection"
	"stabulum/internal/testfixture/mocks"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func NewContainer(cfg config.Config) (*Container, connection.Close, error) {
	panic(
		wire.Build(
			appSet,

			productionDependenciesSet,
		),
	)
}

var appSet = wire.NewSet(
	configSet,

	apiSet,

	appproduct.NewUsecases,
)

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

var productionDependenciesSet = wire.NewSet(
	newContainer,

	productPostgresRepositorySet,
)

var productPostgresRepositorySet = wire.NewSet(
	postgres.NewConnection,

	wire.Bind(new(product.Repository), new(*pgproduct.Repository)),
	pgproduct.NewRepository,
)

func NewTestContainer(cfg config.Config, mockCfg mocks.Config) *TestContainer {
	panic(
		wire.Build(
			appSet,

			testDependenciesSet,
		),
	)
}

var testDependenciesSet = wire.NewSet(
	newTestContainer,
	httpserver.NewTestServer,

	productMockRepositorySet,
)

var productMockRepositorySet = wire.NewSet(
	wire.Bind(new(product.Repository), new(*mockproduct.Repository)),
	mocks.NewProductRepositoryMock,
)
