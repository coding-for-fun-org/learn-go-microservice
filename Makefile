FRONT_END_BINARY=frontendApp
BROKER_BINARY=brokerApp

up:
	docker compose up

up_build:
	docker compose down --remove-orphans
	docker compose up --build

down:
	docker compose down --remove-orphans

build_broker:
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o $(BROKER_BINARY) ./cmd/api

build_frontend:
	cd ./frontend && env GOOS=linux CGO_ENABLED=0 go build -o $(FRONT_END_BINARY) ./cmd/web
