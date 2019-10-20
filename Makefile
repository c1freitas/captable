# Makefile with helpful commands
#

BIN             = captable
OUTPUT_DIR      = build

.PHONY: help 
.DEFAULT_GOAL := help


## Build ##
build: clean ## builds the binary
	go build -o $(OUTPUT_DIR)/$(BIN) .

tests: ## Runs the unit tests
	go test -cover -v ./...

clean: ## Remove build artifacts
	$(RM) $(OUTPUT_DIR)/*

update: ## update go modules
	go mod tidy

help: ## Display this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_\/-]+:.*?## / {printf "\033[34m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | \
		sort | \
		grep -v '#'
