.PHONY: run
run:
	go run cmd/store/*.go

.PHONY: run-env
run-env:
	docker-compose -f deployments/docker-compose.yml up -d

.PHONY: stop-env
stop-env:
	docker-compose -f deployments/docker-compose.yml down -v 

.PHONY: wire-up-modules 
wire-up-modules:
	wire ./internal/infrastructure/di

MOCKERY_CMD=mockery --case=underscore

.PHONY: mocks
mocks:
	${MOCKERY_CMD} --dir=./internal/domain/product --output=./internal/domain/product/mocks --name=Repository
	${MOCKERY_CMD} --dir=./internal/app/queries --output=./internal/app/queries/mocks --name=ProductQuerier