terminal:
	go run ./cmd/terminal

server:
	go run ./cmd/server

build:
	go build -o server ./cmd/server/

docker:
	docker-compose -f ./deployments/docker-compose.yml up --build

all: build