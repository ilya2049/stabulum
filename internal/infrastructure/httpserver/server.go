package httpserver

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"stabulum/internal/common/logger"
	"syscall"
	"time"
)

type Config struct {
	Address string
}

type Server struct {
	httpServer *http.Server
	logger     logger.Logger
}

func New(config Config, handler http.Handler, logger logger.Logger) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: config.Address,

			Handler: handler,
		},

		logger: logger,
	}
}

func (s *Server) ListenAndServeAsync() {
	s.logger.Println("http server is ready to accept connections")

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				s.logger.Println("failed to listen and serve an http server:", err.Error())
			}
		}
	}()
}

func (s *Server) Shutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Println("failed to shutdown an http server:", err.Error())

		return
	}

	s.logger.Println("http server successfully stopped")
}
