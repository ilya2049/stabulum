package httpserver

import "net/http"

type Config struct {
	Adderss string
}

type Server struct {
	*http.Server
}

func New(config Config, handler http.Handler) *Server {
	return &Server{
		Server: &http.Server{
			Addr: config.Adderss,

			Handler: handler,
		},
	}
}
