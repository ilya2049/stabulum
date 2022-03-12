//go:build wireinject
// +build wireinject

package di

import (
	"net/http"

	appproduct "stabulum/internal/app/product"
	"stabulum/internal/common/logger"
	"stabulum/internal/common/testfixture"
	"stabulum/internal/common/testfixture/mocks"
	"stabulum/internal/domain/product"
	mockproduct "stabulum/internal/domain/product/mocks"
	"stabulum/internal/infrastructure/api/router"
	apiproduct "stabulum/internal/infrastructure/api/router/product"
	"stabulum/internal/infrastructure/config"
	"stabulum/internal/infrastructure/httpserver"
	infrastructureLogger "stabulum/internal/infrastructure/logger"
	"stabulum/internal/infrastructure/postgres"
	pgproduct "stabulum/internal/infrastructure/postgres/product"
	"stabulum/internal/pkg/connection"

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
	loggerSet,

	newContainer,

	productPostgresRepositorySet,
)

var loggerSet = wire.NewSet(
	wire.Bind(new(logger.Logger), new(*infrastructureLogger.Logger)),
	infrastructureLogger.New,
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
	spyLoggerSet,

	newTestContainer,
	httpserver.NewTestServer,

	productMockRepositorySet,
)

var spyLoggerSet = wire.NewSet(
	wire.Bind(new(logger.Logger), new(*testfixture.SpyLogger)),
	testfixture.NewSpyLogger,
)

var productMockRepositorySet = wire.NewSet(
	wire.Bind(new(product.Repository), new(*mockproduct.Repository)),
	mocks.NewProductRepositoryMock,
)
