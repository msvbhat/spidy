GO_FILES := $(shell find . -type f -name '*.go' | grep -v /vendor/ | grep -v _test.go)
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
BINARY := spidy

.PHONY: all dep lint build clean

all: build

dep: ## Get the dependencies
	@go get -v -d ./...

lint: ## Run the golint
	@golint -set_exit_status ${PKG_LIST}

build: ## Build the static binary
	@env GOARCH=amd64 GOOS=darwin go build -i -o ${BINARY}

clean: ## Remove the previous build
	@rm -f ${BINARY}
