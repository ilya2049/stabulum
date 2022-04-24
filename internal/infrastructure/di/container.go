package di

import (
	appproduct "stabulum/internal/app/product"
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
	loggerLogger := logger.New()
	postgresConfig := config.NewPostgresConfig(cfg)
	db, closePostgresConnection, err := postgres.NewConnection(postgresConfig, loggerLogger)
	if err != nil {
		return nil, nil, err
	}
	repository := pgproduct.NewRepository(db)
	usecases := appproduct.NewUsecases(repository)
	querier := queries.NewQuerier(db, loggerLogger)
	handler := apiroduct.NewHandler(usecases, querier)
	engine := router.New(handler)
	server := httpserver.New(httpserverConfig, engine, loggerLogger)
	container := newContainer(server)
	return container, func() {
		closePostgresConnection()
	}, nil
}
