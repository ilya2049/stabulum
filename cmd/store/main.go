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

	defer closeConnections()

	server := diContainer.APIHTTPServer
	defer server.Shutdown()
	server.ListenAndServeAsync()
}
