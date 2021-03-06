APP_NAME=pubsub-app

.PHONY: check-%
check-%: ## check environment variable is exists
	@if [ -z '${${*}}' ]; then echo 'Environment variable $* not set' && exit 1; fi

.PHONY: help
help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean:  ## remove artifacts
	@rm -rf coverage.txt profile.out ./bin
	@echo Successfuly removed artifacts

.PHONY: test-unit
test-unit: ## execute unit test
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: lint
lint: ## execute golint
	@golint ./...

.PHONY: report
report: ## execute goreportcard
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/$(APP_NAME)'

.PHONY: up-compose
up-compose: ## run docker-compose up
	@docker-compose -p $(APP_NAME) -f ./deployments/docker-compose.yml up -d

.PHONY: down-compose
down-compose: ## run docker-compose down
	@docker-compose -p $(APP_NAME) -f ./deployments/docker-compose.yml down -v
