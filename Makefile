## ----------------------------------------------------------------------
## The purpose of this Makefile is to simplify common development tasks.
## ----------------------------------------------------------------------
##

.PHONY:terminal
terminal: ## run the app in terminal mode
##
	go run ./cmd/terminal

.PHONY:server
server: ## run the app in server mode
##
	go run ./cmd/server

.PHONY:build_server
build_server: ## build binary for server mode
##
	go build -o server ./cmd/server/

.PHONY:docker
docker: ## run in docker container
##
	docker-compose -f ./deployments/docker-compose.yml up --build

.PHONY:help
help: ## Show this help
##
	@sed -ne '/@sed/!s/##//p' $(MAKEFILE_LIST)