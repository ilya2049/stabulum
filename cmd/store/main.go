package main

import (
	"log"
	"stabulum/internal/infrastructure/config"
	"stabulum/internal/infrastructure/di"
)

func main() {
	diContainer := di.NewTestContainer(
		config.ReadFromMemory(),
		config.ReadFromMemoryMockConfig(),
	)

	server := diContainer.APIHTTPServer

	log.Fatal(server.ListenAndServe())
}
