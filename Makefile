.PHONY: build test lint clean install

BINARY := mkexp
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -X github.com/dreikanter/mkexp/cmd.Version=$(VERSION)

build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) .

test:
	go test ./...

lint:
	go tool golangci-lint run

clean:
	rm -f $(BINARY)

install:
	go install -ldflags "$(LDFLAGS)" .
