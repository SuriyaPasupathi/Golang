# Makefile for building, testing, and running the project

build:
	@echo "Building services..."
	@go build -o bin/event-service ./cmd/event-service
	@go build -o bin/booking-service ./cmd/booking-service
	@go build -o bin/gateway-service ./cmd/gateway-service

proto:
	@echo "Generating protobuf code..."
	@./scripts/proto-gen.sh

run:
	@echo "Running services..."
	@docker-compose -f ./deployments/docker/docker-compose.yaml up --build

deploy:
	@echo "Deploying to Docker Swarm..."
	@./scripts/deploy.sh

clean:
	@echo "Cleaning up binaries..."
	@rm -rf bin/*
