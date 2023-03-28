BINARY_NAME=projector-controller

all: test build

test:
	go test -race -v ./...

build:
	goreleaser build --snapshot --clean

clean:
	rm -rf dist

.PHONY: all
