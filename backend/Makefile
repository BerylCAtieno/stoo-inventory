.PHONY: build run dev test clean

BINARY_NAME := stoo-inventory-app

PROJECT_ROOT := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

BUILD_DIR := ./bin

MAIN_PACKAGE := ./cmd/server

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"


run: build
	@echo "Running $(BINARY_NAME) from root..."
	cd $(PROJECT_ROOT) && $(BUILD_DIR)/$(BINARY_NAME)


dev:
	@echo "Starting development server with Air..."
	air


test:
	@echo "Running tests..."
	go test ./... -v


clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."
