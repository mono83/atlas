# Makefile configuration
.DEFAULT_GOAL := help

deps: ## Downloads dependencies
	go get github.com/go-sql-driver/mysql
	go get github.com/spf13/cobra
	go get github.com/mitchellh/go-homedir
	go get github.com/go-ini/ini
	go get github.com/stretchr/testify/assert
	go get github.com/fzipp/gocyclo

test: ## Runs unit tests
	go test ./...

vet: ## Code check
	gofmt -s -w .
	go vet ./...
	gocyclo -over 17 .
	golint ./...

travis: deps vet test ## Runs all tasks for travis CI

help:
	@grep --extended-regexp '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
