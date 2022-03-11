package httpserver

import "net/http"

type Config struct {
	Address string
}

type Server struct {
	*http.Server
}

func New(config Config, handler http.Handler) *Server {
	return &Server{
		Server: &http.Server{
			Addr: config.Address,

			Handler: handler,
		},
	}
}
