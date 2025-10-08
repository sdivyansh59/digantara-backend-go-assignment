.PHONY: wire run build clean test help

# Generate wire dependency injection code
wire:
	@echo "Generating wire dependencies..."
	@wire ./app

# Run the application
run:
	@echo "Running application..."
	@go run main.go

# Build the application
build:
	@echo "Building application..."
	@go build -o bin/api-service main.go

# Clean generated files
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f app/wire_gen.go

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Help command
help:
	@echo "Available commands:"
	@echo "  make wire    - Generate wire dependency injection code"
	@echo "  make run     - Run the application"
	@echo "  make build   - Build the application binary"
	@echo "  make clean   - Clean generated files and binaries"
	@echo "  make test    - Run tests"
	@echo "  make deps    - Install and tidy dependencies"
	@echo "  make help    - Show this help message"

