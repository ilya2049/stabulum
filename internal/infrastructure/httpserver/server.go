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

	quitChannel chan os.Signal
}

func New(config Config, handler http.Handler, logger logger.Logger) *Server {
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	return &Server{
		httpServer: &http.Server{
			Addr: config.Address,

			Handler: handler,
		},

		logger:      logger,
		quitChannel: quitChannel,
	}
}

func (s *Server) ListenAndServeAsync() {
	s.logger.Println("http server is ready to accept connections")

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				s.logger.Println("failed to listen and serve an http server:", err.Error())
				s.quitChannel <- syscall.SIGINT
			}
		}
	}()
}

func (s *Server) WaitForShutdown() {
	<-s.quitChannel

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Println("failed to shutdown an http server:", err.Error())

		return
	}

	s.logger.Println("http server successfully stopped")
}
