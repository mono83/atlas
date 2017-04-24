# Makefile configuration
.DEFAULT_GOAL := help

deps: ## Downloads dependencies
	go get github.com/go-sql-driver/mysql
	go get github.com/spf13/cobra

test: ## Runs unit tests
	go test ./...

travis: deps test ## Runs all tasks for travis CI

help:
	@grep --extended-regexp '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
