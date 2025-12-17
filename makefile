.PHONY: build run

BINARY_NAME=promo-bots
BUILD_DIR=./bin

export CONFIG_PATH ?= ./configs

build:
	@echo "Building $(BINARY_NAME)"
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/app

run: build
	@echo "Running $(BINARY_NAME)"
	@$(BUILD_DIR)/$(BINARY_NAME)