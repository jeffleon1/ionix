.PHONY: all build run clean

APP_NAME := ionix
SRC_DIR := ./cmd


run-docker:
	@echo "Running $(APP_NAME) in Docker..."
	@docker compose up -d --build

down-docker:
	@echo "Stopping $(APP_NAME)..."
	@docker compose down

docs:
	@echo "Making docs $(APP_NAME)..."
	@swag init --parseDependency --parseInternal -g $(SRC_DIR)/main.go