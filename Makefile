# Variables
APP_NAME := validate
PKG := ./...
LINTER := golangci-lint run

# Default target
all: build test lint

# Build the Go application
build:
	go build -o $(APP_NAME) ./main.go

# Run tests
test:
	go test $(PKG)

lint:
	$(LINTER)

# Clean up build artifacts
clean:
	rm -f $(APP_NAME)

# Run the application (you can pass arguments as needed)
run: build
	./$(APP_NAME)

# Phony targets (to prevent conflicts with file names)
.PHONY: all build test lint clean run
