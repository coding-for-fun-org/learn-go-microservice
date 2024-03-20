FRONT_END_BINARY=frontendApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
LOGGER_BINARY=loggerApp

up:
	docker compose up

up_build: build_frontend build_broker build_auth build_logger
	docker compose down --remove-orphans
	docker compose up --build

down:
	docker compose down --remove-orphans

purge:
	docker compose down --remove-orphans -v --rmi all

soft-purge:
	docker compose down --remove-orphans --rmi all

build_frontend:
	cd ./frontend && env GOOS=linux CGO_ENABLED=0 go build -o $(FRONT_END_BINARY) ./cmd/web

build_broker:
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o $(BROKER_BINARY) ./cmd/api

build_auth:
	cd ./auth-service && env GOOS=linux CGO_ENABLED=0 go build -o $(AUTH_BINARY) ./cmd/api

build_logger:
	cd ./logger-service && env GOOS=linux CGO_ENABLED=0 go build -o $(LOGGER_BINARY) ./cmd/api
