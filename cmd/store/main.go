package main

import (
	"stabulum/internal/infrastructure/config"
	"stabulum/internal/infrastructure/di"
)

func main() {
	diContainer, closeConnections, err := di.NewContainer(
		config.ReadFromMemory(),
	)

	logger := diContainer.Logger

	if err != nil {
		logger.Println(err)

		return
	}

	defer closeConnections()

	server := diContainer.APIHTTPServer

	logger.Println(server.ListenAndServe())
}
