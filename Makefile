BINARY_NAME=projector-controller

all: check test build

check:
ifeq (, $(shell which golangci-lint))
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin latest
endif
	golangci-lint run
	go mod tidy

test:
	go test -race -v ./...

build:
	goreleaser build --snapshot --clean

clean:
	rm -rf dist

.PHONY: all
