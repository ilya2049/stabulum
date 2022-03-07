package di

import (
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
