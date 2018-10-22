all: build

help: ## Show usage
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build binaries
	@echo 'Building binary...'
	@go build -o bin/parser

run: ## Run parser
	@echo 'Starting parser...'
	@echo '' > ./logs/stdout.log
	@bin/parser --conn="root:daniyar@/mitdb?parseTime=true&loc=Local" --stdout=./logs/stdout.log --limit=1

dump: ## Dump database models
	@sqlboiler -o pkg/models -p models --wipe mysql

test-models: ## Test generated models
	bash -c "cd pkg/models && go test --test.config=../../sqlboiler.yml ./pkg/models && cd ../../"
