# Makefile for the Golang MFE Server project with main.go in cmd/server

# Define Go related variables
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOCMD := $(GOBASE)/cmd/app

# Define Docker related variables
DOCKER_IMG_NAME := mfe-server
DOCKER_CONTAINER_NAME := mfe-server

.PHONY: all build run clean docker-build docker-run docker-clean test

all: build

# Build the project
build:
	@echo "Building the project..."
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/server $(GOCMD)

# Run the project
run: build
	@echo "Running the server..."
	@$(GOBIN)/server

# Clean up
clean:
	@echo "Cleaning up..."
	@go clean
	@rm -rf $(GOBIN)/server

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMG_NAME) .

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	@docker run --rm -d --name $(DOCKER_CONTAINER_NAME) -p 8000:8000 $(DOCKER_IMG_NAME)

# Stop and remove Docker container
docker-clean:
	@echo "Stopping and removing Docker container..."
	@docker stop $(DOCKER_CONTAINER_NAME)
	@docker rm $(DOCKER_CONTAINER_NAME)

# Run tests
test:
	@echo "Running tests..."
	@go test ./...
