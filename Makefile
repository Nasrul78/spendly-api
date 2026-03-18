.PHONY: run build migrate-up migrate-down sqlc swagger fmt lint

run:
	air

build:
	go build -o ./tmp/main ./cmd/api

migrate-up:
	goose up

migrate-down:
	goose down

sqlc:
	sqlc generate

swagger:
	swag init -g cmd/api/main.go --output docs

fmt:
	gofmt -w .
	goimports -w .
	swag fmt

lint:
	go vet ./...
	golangci-lint run ./...
	golangci-lint fmt
