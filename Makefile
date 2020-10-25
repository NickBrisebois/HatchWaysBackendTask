GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=hwApi

all: build

.PHONY: test build
build:
	rm -rf ./build/;
	mkdir ./build;
	cp -r ./config.toml ./build/
	$(MAKE) -s go-build

go-build:
	@GOPATH=$(GOPATH) go build -o ./build/$(BINARY_NAME)

