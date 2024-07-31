## ----------------------------------------------------------------------
## The purpose of this Makefile is to simplify common development tasks.
## ----------------------------------------------------------------------
##
## Usage:
##   - make build        : Build the application
##   - make terminal     : Run the application in terminal mode
##   - make help         : Show available commands and descriptions
##


PROG = $(shell basename `git rev-parse --show-toplevel`)

.PHONY:build
build:
	go build -o $(PROG)

.PHONY:terminal
terminal:
	go run . -mode=terminal

.PHONY:help
help:
	@sed -ne '/@sed/!s/##//p' $(MAKEFILE_LIST)