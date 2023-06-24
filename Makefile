export GO111MODULE=on

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: test
test:
	gotestsum --format=testname -- -race -tags=units,integrations --cover ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: update
update:
	go mod tidy
	go mod verify

.PHONY: generate-openapi
generate-openapi:
	swag init --pd 2
	redoc-cli bundle -o docs/doc.html docs/swagger.yaml
	gzip --force docs/doc.html

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'