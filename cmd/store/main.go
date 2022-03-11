package main

import (
	"log"
	"stabulum/internal/infrastructure/config"
	"stabulum/internal/infrastructure/di"
)

func main() {
	diContainer, closeConnections, err := di.NewTestContainer(
		config.ReadFromMemory(),
		config.ReadFromMemoryMockConfig(),
	)

	if err != nil {
		log.Println(err)

		return
	}

	defer closeConnections()

	server := diContainer.APIHTTPServer

	log.Println(server.ListenAndServe())
}
