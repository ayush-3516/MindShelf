.PHONY: build run test clean docker-dev docker-prod docker-clean docker-logs certs help

# Default target
help:
	@echo "Mindshelf Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build       Build the Go application"
	@echo "  make run         Run the application locally"
	@echo "  make test        Run tests"
	@echo "  make clean       Clean build artifacts"
	@echo ""
	@echo "Docker commands:"
	@echo "  make docker-dev  Start development environment"
	@echo "  make docker-prod Start production environment"
	@echo "  make docker-clean Clean Docker resources"
	@echo "  make docker-logs View Docker logs"
	@echo "  make certs      Generate SSL certificates"

# Build the Go application
build:
	go build -o bin/mindshelf-api ./cmd/api

# Run the application
run: build
	./bin/mindshelf-api

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Docker development environment
docker-dev:
	@bash ./mindshelf.sh start

# Docker production environment
docker-prod:
	@bash ./deploy-prod.sh

# Clean Docker resources
docker-clean:
	@bash ./mindshelf.sh clean

# View Docker logs
docker-logs:
	@bash ./mindshelf.sh logs

# Generate certificates
certs:
	@bash ./generate-certs.sh