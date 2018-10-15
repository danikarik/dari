all: build

help: ## Show usage
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build binaries
	@echo 'Building binary...'
	@go build -o bin/parser

run: ## Run parser
	@echo 'Starting parser...'
	@bin/parser
