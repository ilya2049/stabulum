package di

import (
	appproduct "stabulum/internal/app/product"
	"stabulum/internal/common/event"
	"stabulum/internal/infrastructure/api/router"
	apiroduct "stabulum/internal/infrastructure/api/router/product"
	"stabulum/internal/infrastructure/config"
	"stabulum/internal/infrastructure/httpserver"
	"stabulum/internal/infrastructure/logger"
	"stabulum/internal/infrastructure/postgres"
	pgproduct "stabulum/internal/infrastructure/postgres/product"
	"stabulum/internal/infrastructure/postgres/queries"
)

type Container struct {
	APIHTTPServer *httpserver.Server
}

func newContainer(apiHTTPServer *httpserver.Server) *Container {
	return &Container{
		APIHTTPServer: apiHTTPServer,
	}
}

func NewContainer(cfg config.Config) (*Container, func(), error) {
	httpserverConfig := config.NewHTTPServerConfig(cfg)
	aLogger := logger.New()
	loggerEventHandler := logger.NewEventHandler(aLogger)
	postgresConfig := config.NewPostgresConfig(cfg)
	postgresConnection, closePostgresConnection, err := postgres.NewConnection(postgresConfig, aLogger)
	if err != nil {
		return nil, nil, err
	}

	eventBus := event.NewBus()
	repository := pgproduct.NewRepository(postgresConnection.SQLDB, postgresConnection.SQLDB, eventBus)
	loggerEventHandler.RegisterEvents(eventBus)
	usecases := appproduct.NewUsecases(repository)
	querier := queries.NewQuerier(postgresConnection.GormDB, aLogger)
	handler := apiroduct.NewHandler(usecases, querier)
	engine := router.New(handler)
	server := httpserver.New(httpserverConfig, engine, aLogger)
	container := newContainer(server)
	return container, func() {
		closePostgresConnection()
	}, nil
}
