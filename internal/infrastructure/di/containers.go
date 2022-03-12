package di

import (
	"net/http/httptest"
	"stabulum/internal/common/logger"
	"stabulum/internal/infrastructure/httpserver"
)

type Container struct {
	APIHTTPServer *httpserver.Server
	Logger        logger.Logger
}

func newContainer(apiHTTPServer *httpserver.Server, logger logger.Logger) *Container {
	return &Container{
		APIHTTPServer: apiHTTPServer,
		Logger:        logger,
	}
}

type TestContainer struct {
	APIHTTPTestServer *httptest.Server
}

func newTestContainer(apiHTTPTestServer *httptest.Server) *TestContainer {
	return &TestContainer{
		APIHTTPTestServer: apiHTTPTestServer,
	}
}
