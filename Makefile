# Makefile for Spoticus: AWS Spot Cluster Bot

BINARY_NAME := spoticus
CMD_PATH := cmd/$(BINARY_NAME)/main.go
OUTPUT_DIR := bin
OUTPUT_BIN := $(OUTPUT_DIR)/$(BINARY_NAME)

.PHONY: all build run fmt vet lint clean help

# Default target
all: build

# Build the bot binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(OUTPUT_DIR)
	go build -o $(OUTPUT_BIN) $(CMD_PATH)

# Run the bot locally
run:
	@echo "Running $(BINARY_NAME)..."
	go run $(CMD_PATH)

# Format Go code
fmt:
	@echo "Formatting Go code..."
	go fmt ./...

# Vet for potential issues
vet:
	@echo "Running go vet..."
	go vet ./...

# Lint the codebase
lint:
	@echo "Running golangci-lint..."
	golangci-lint run

# Clean up build artifacts
clean:
	@echo "Cleaning build output..."
	rm -rf $(OUTPUT_DIR)

# Show help
help:
	@echo ""
	@echo "Spoticus Bot – Makefile Help"
	@echo ""
	@echo "Targets:"
	@echo "  make build     - Build the binary to ./bin/$(BINARY_NAME)"
	@echo "  make run       - Run the bot locally"
	@echo "  make fmt       - Format Go code"
	@echo "  make vet       - Static analysis with go vet"
	@echo "  make
