package di

import (
	"net/http/httptest"
	"stabulum/internal/common/testfixture"
	"stabulum/internal/infrastructure/httpserver"
)

type Container struct {
	APIHTTPServer *httpserver.Server
}

func newContainer(apiHTTPServer *httpserver.Server) *Container {
	return &Container{
		APIHTTPServer: apiHTTPServer,
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
