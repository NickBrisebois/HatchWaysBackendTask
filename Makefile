GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=hwApi

all: test build

.PHONY: test build
build:
	go mod vendor;
	rm -rf ./build/;
	mkdir ./build;
	cp -r ./config.toml ./build/
	$(MAKE) -s go-build

go-build:
	@GOPATH=$(GOPATH) $(GOBUILD) -o ./build/$(BINARY_NAME)

test:
	$(GOTEST) -v ./...
