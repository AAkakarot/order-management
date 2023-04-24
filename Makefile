.PHONY: build run test lint clean

build:
	go build -o ORDER-MANAGEMENT cmd/main.go

run:
	go run cmd/main.go

test:
	go test -v ./...

lint:
	golangci-lint run