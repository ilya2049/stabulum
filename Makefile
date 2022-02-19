.PHONY: run
run:
	go run cmd/store/*.go

.PHONY: wire-up-modules 
wire-up-modules:
	wire ./internal/infrastructure/di