.PHONY: run
run:
	go run cmd/store/*.go

.PHONY: run-env
run-env:
	docker-compose -f deployments/docker-compose.yml up -d

.PHONY: stop-env
stop-env:
	docker-compose -f deployments/docker-compose.yml down -v 