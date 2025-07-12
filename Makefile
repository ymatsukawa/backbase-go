APP_NAME := backbase-go
BUILD_DIR := ./bin
MAIN_PATH := ./cmd/api
DOCKER_COMPOSE := docker-compose.yml


build:
	@echo "Building $(APP_NAME)..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

test:
	@echo "Running tests..."
	go test -v ./...

coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

up:
	@echo "dev env up"
	docker compose up

down:
	@echo "dev env down"
	docker compose down

debug:
	@echo "dev env debug"
	DEBUG_MODE=true docker compose up

new-migration:
	@echo "Creating new database table migration..."
	@read -p "Enter migration name: " name; \
	if [ -z "$$name" ]; then \
		echo "Migration name cannot be empty"; \
		exit 1; \
	fi; \
	docker exec api sql-migrate new $$name

migrate:
	@echo "Running database migrations..."
	docker exec api sql-migrate up $(name)

rollback:
	@echo "Rolling back database migrations..."
	docker exec api sql-migrate down

gen-entity:
	@echo "Generating entities..."
	docker exec api sqlc generate
