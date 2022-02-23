.PHONY: run
run:
	go run cmd/store/*.go

.PHONY: wire-up-modules 
wire-up-modules:
	wire ./internal/infrastructure/di

MOCKERY_CMD=mockery --keeptree --case=underscore

.PHONY: mocks
mocks:
	${MOCKERY_CMD} --dir=./internal/domain/product --output=./internal/domain/product/mocks --name=Repository