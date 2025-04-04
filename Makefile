.PHONY: install-tools
install-tools:
	@./scripts/install-tools.sh

.PHONY: run-app
run-app:
	@./scripts/run.sh  

.PHONY: build
build:
	@./scripts/build.sh 

.PHONY: install-dependencies
install-dependencies:
	@./scripts/install-dependencies.sh 

.PHONY: docker-compose-infra-up
docker-compose-infra-up:
	@docker-compose -f deployments/docker-compose/docker-compose.infrastructure.yaml up --build -d

docker-compose-infra-down:
	@docker-compose -f deployments/docker-compose/docker-compose.infrastructure.yaml down

.PHONY: openapi
openapi:
	@./scripts/openapi.sh 

.PHONY: unit-test
unit-test:
	@./scripts/test.sh unit

.PHONY: integration-test
integration-test:
	@./scripts/test.sh integration

.PHONY: e2e-test
e2e-test:
	@./scripts/test.sh e2e

#.PHONY: load-test
#load-test:
#	@./scripts/test.sh load-test

.PHONY: format
format:
	@./scripts/format.sh 

.PHONY: lint
lint:
	@./scripts/lint.sh 

.PHONY: mocks
mocks:
	mockery --output mocks --all

.PHONY: update-dependencies
update-dependencies:
	@./scripts/update-dependencies.sh
