.PHONY: all build test lint clean

all: test build

build:
	go build -ldflags="-s -w" -o bin/loglinter ./cmd/loglinter
	go build -buildmode=plugin -o bin/loglinter.so ./plugin

test:
	go test -v -race -count=1 ./pkg/linter/...

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/
