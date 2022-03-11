package di

import (
	"net/http/httptest"
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
}

func newTestContainer(apiHTTPTestServer *httptest.Server) *TestContainer {
	return &TestContainer{
		APIHTTPTestServer: apiHTTPTestServer,
	}
}
