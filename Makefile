.DEFAULT_GOAL := help
.PHONY: pretty build clean help

EXECUTABLE=superperm
OS=darwin
ARCH=arm64

pretty: ## Format golang code
	@goimports -w cmd internal sdk

build: pretty ## Build binary for specific OS and ARCH
	@go build -ldflags='-s -w' -o bin/$(EXECUTABLE)_$(OS)_$(ARCH) cmd/$(EXECUTABLE)/main.go
	@echo bin/$(EXECUTABLE)_$(OS)_$(ARCH)

build-all: ## Build for all platforms
	@make OS=darwin ARCH=amd64 build
	@make OS=linux  ARCH=amd64 build
	@make OS=linux  ARCH=arm64 build

clean: ## Delete bin folder
	@rm -rf bin

help: ## Help menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
