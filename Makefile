terminal:
	go run ./cmd/terminal

server:
	go run ./cmd/server

build:
	go build -o server ./cmd/server/

all: build