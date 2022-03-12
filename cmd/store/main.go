package main

import (
	"log"
	"stabulum/internal/infrastructure/config"
	"stabulum/internal/infrastructure/di"
)

func main() {
	diContainer, closeConnections, err := di.NewContainer(
		config.ReadFromMemory(),
	)

	if err != nil {
		log.Println(err)

		return
	}

	logger := diContainer.Logger

	defer closeConnections()

	server := diContainer.APIHTTPServer

	logger.Println(server.ListenAndServe())
}
