.PHONY: migrate-up migrate-down migrate-reset migrate-status migrate-create

BINARY_NAME=promo-bots
BUILD_DIR=./bin
MIGRATIONS_DIR := ./migrations


# Colors for output
COLOR_RESET=\033[0m
COLOR_BOLD=\033[1m
COLOR_GREEN=\033[32m
COLOR_YELLOW=\033[33m

export CONFIG_PATH ?= ./configs

build:
	@echo "Building $(BINARY_NAME)"
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/app

run: build
	@echo "Running $(BINARY_NAME)"
	@$(BUILD_DIR)/$(BINARY_NAME)
	

migrate-up: ## Применить все миграции
	@echo "$(COLOR_YELLOW)Running migrations...$(COLOR_RESET)"
	@if [ -z "$(DATABASE_URL)" ]; then \
		echo "ERROR: DATABASE_URL is not set"; \
		exit 1; \
	fi
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DATABASE_URL) goose -dir $(MIGRATIONS_DIR) up
	@echo "$(COLOR_GREEN)Migrations applied!$(COLOR_RESET)"

migrate-down: ## Откатить последнюю миграцию
	@echo "$(COLOR_YELLOW)Rolling back last migration...$(COLOR_RESET)"
	@if [ -z "$(DATABASE_URL)" ]; then \
		echo "ERROR: DATABASE_URL is not set"; \
		exit 1; \
	fi
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DATABASE_URL) goose -dir $(MIGRATIONS_DIR) down
	@echo "$(COLOR_GREEN)Migration rolled back!$(COLOR_RESET)"

migrate-reset: ## Сбросить все миграции
	@echo "$(COLOR_YELLOW)Resetting all migrations...$(COLOR_RESET)"
	@if [ -z "$(DATABASE_URL)" ]; then \
		echo "$(COLOR_YELLOW)ERROR: DATABASE_URL is not set$(COLOR_RESET)"; \
		exit 1; \
	fi
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DATABASE_URL) goose -dir $(MIGRATIONS_DIR) reset
	@echo "$(COLOR_GREEN)All migrations reset!$(COLOR_RESET)"

migrate-status: ## Показать статус миграций
	@if [ -z "$(DATABASE_URL)" ]; then \
		echo "ERROR: DATABASE_URL is not set"; \
		exit 1; \
	fi
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DATABASE_URL) goose -dir $(MIGRATIONS_DIR) status

migrate-create: ## Создать новую миграцию (use: make migrate-create NAME=migration_name)
	@if [ -z "$(name)" ]; then \
		echo -e "$(COLOR_YELLOW)ERROR: migration name is required$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo -e "$(COLOR_YELLOW)Creating new migration: $(NAME)...$(COLOR_RESET)"
	@goose -dir $(MIGRATIONS_DIR) create $(name) sql
	@echo -e "$(COLOR_GREEN)Migration created!$(COLOR_RESET)

sqlc-gen: ## Generate code from SQL
	@echo "$(COLOR_YELLOW)Generating code from SQL...$(COLOR_RESET)"
	sqlc generate
	@echo "$(COLOR_GREEN)Code generated!$(COLOR_RESET)"
