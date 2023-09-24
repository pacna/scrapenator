## ----------------------------------------------------------------------
## The purpose of this Makefile is to simplify common development tasks.
## ----------------------------------------------------------------------
##

PROG = $(shell basename `git rev-parse --show-toplevel`)

.PHONY:build
build: ## Build the application 
##
	go build -o $(PROG)

.PHONY:terminal
terminal: ## Run the application in terminal mode
##
	go run . -mode=terminal

.PHONY:server
server: ## Run the application in server mode
##
	go run . -mode=server

.PHONY:docker
docker: ## Run the application in a Docker container
##
	docker-compose -f ./deployments/docker-compose.yml up --build

.PHONY:help
help: ## Show the help message with target descriptions
##
	@sed -ne '/@sed/!s/##//p' $(MAKEFILE_LIST)