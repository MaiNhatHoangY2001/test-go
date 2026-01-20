.PHONY: help build run test test-coverage test-unit test-integration clean docker-build docker-up docker-down deps fmt lint

# Variables
BINARY_NAME=server
BINARY_PATH=./cmd/server/main.go
DOCKER_IMAGE=todo-app
DOCKER_TAG=latest

help:
	@echo "Available commands:"
	@echo "  make build              - Build the application"
	@echo "  make run                - Run the application"
	@echo "  make test               - Run all tests"
	@echo "  make test-unit          - Run unit tests only"
	@echo "  make test-integration   - Run integration tests only"
	@echo "  make test-coverage      - Run tests with coverage report"
	@echo "  make clean              - Clean build artifacts"
	@echo "  make deps               - Download dependencies"
	@echo "  make fmt                - Format code"
	@echo "  make lint               - Run linter"
	@echo "  make docker-build       - Build Docker image"
	@echo "  make docker-up          - Start Docker containers"
	@echo "  make docker-down        - Stop Docker containers"

# Dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Build
build:
	@echo "Building application..."
	go build -o $(BINARY_NAME) $(BINARY_PATH)
	@echo "Build successful: $(BINARY_NAME)"

# Run
run: build
	@echo "Running application..."
	./$(BINARY_NAME)

# Testing
test: test-unit test-integration
	@echo "All tests completed!"

test-unit:
	@echo "Running unit tests..."
	go test -v -race -coverprofile=coverage.out ./tests/unit/...
	@echo "Unit tests completed!"

test-integration:
	@echo "Running integration tests..."
	go test -v -race ./tests/integration/...
	@echo "Integration tests completed!"

test-coverage: test-unit
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-coverage-check:
	@echo "Checking coverage percentage..."
	@go test -v -race -coverprofile=coverage.out ./tests/unit/... ./tests/integration/...
	@go tool cover -func=coverage.out | grep total | awk '{print "Total coverage: " $$3}'

# Code Quality
fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "Code formatted!"

lint:
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...
	@echo "Linting completed!"

vet:
	@echo "Running go vet..."
	go vet ./...
	@echo "Go vet completed!"

# Docker
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	@echo "Docker image built: $(DOCKER_IMAGE):$(DOCKER_TAG)"

docker-up:
	@echo "Starting Docker containers..."
	docker-compose up -d
	@echo "Docker containers started!"

docker-down:
	@echo "Stopping Docker containers..."
	docker-compose down
	@echo "Docker containers stopped!"

docker-logs:
	docker-compose logs -f

# Cleanup
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html
	@echo "Cleanup completed!"

# Database
migrate-up:
	@echo "Running migrations..."
	go run ./cmd/migrate/main.go up

migrate-down:
	@echo "Rolling back migrations..."
	go run ./cmd/migrate/main.go down

# Development
dev: fmt lint test build run

# Full pipeline
all: clean deps fmt lint test build
	@echo "Full pipeline completed!"
