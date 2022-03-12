package di

import (
	"net/http/httptest"
	"stabulum/internal/common/logger"
	"stabulum/internal/common/testfixture"
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
	SpyLogger         *testfixture.SpyLogger
}

func newTestContainer(apiHTTPTestServer *httptest.Server, spyLogger *testfixture.SpyLogger) *TestContainer {
	return &TestContainer{
		SpyLogger:         spyLogger,
		APIHTTPTestServer: apiHTTPTestServer,
	}
}
