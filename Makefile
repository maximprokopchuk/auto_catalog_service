.PHONY: build

build: 
		go build -v ./cmd/auto_reference_catalog_service

.PHONY: run

run: 
		go run -v ./cmd/auto_reference_catalog_service

.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.PHONY: proto_generate
proto_generate:
		protoc -I proto --go_out=plugins=grpc:pkg/api proto/auto_reference_catalog.proto

.PHONY: sqlc_generate
sqlc_generate:
		sqlc generate

.PHONY: deps
deps:
		go mod tidy

.PHONY: lint
lint:
		golangci-lint run


DEFAULT_GOAL := build