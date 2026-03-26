.PHONY: help install start stop clean test build restart logs backend-dev frontend-dev test-backend test-frontend lint-backend lint-frontend build-backend build-frontend seed migrate db-shell status

GREEN=\033[0;32m
NC=\033[0m

help: ## Show help
	@echo "$(GREEN)Diputat Fact-Check Platform$(NC)"
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(NC) %s\n", $$1, $$2}'

install: ## Install backend and frontend dependencies
	@echo "$(GREEN)Installing dependencies...$(NC)"
	cd backend && go mod download
	cd frontend && npm install

start: ## Start all services with Docker
	@echo "$(GREEN)Starting services...$(NC)"
	docker-compose up -d --build
	@echo "$(GREEN)Services started:$(NC)"
	@echo "  - Backend: http://localhost:8080"
	@echo "  - Frontend: http://localhost:4200"
	@echo "  - MongoDB: mongodb://localhost:27017"
	@echo "  - Mongo Express: http://localhost:8081"

stop: ## Stop all services
	@echo "$(GREEN)Stopping services...$(NC)"
	docker-compose down

clean: ## Clean containers and local build output
	@echo "$(GREEN)Cleaning...$(NC)"
	docker-compose down -v
	rm -rf backend/vendor frontend/node_modules frontend/dist

restart: stop start ## Restart all services

logs: ## View Docker logs
	docker-compose logs -f

backend-dev: ## Run the backend locally
	@echo "$(GREEN)Starting backend...$(NC)"
	cd backend && go run ./cmd/server

frontend-dev: ## Run the frontend locally
	@echo "$(GREEN)Starting frontend...$(NC)"
	cd frontend && npm run dev -- --host 0.0.0.0 --port 4200

test-backend: ## Run backend tests
	@echo "$(GREEN)Running backend tests...$(NC)"
	cd backend && go test ./...

test-frontend: ## Run frontend checks
	@echo "$(GREEN)Running frontend checks...$(NC)"
	cd frontend && npm test

lint-backend: ## Check backend formatting and static analysis
	@echo "$(GREEN)Linting backend...$(NC)"
	cd backend && test -z "`gofmt -l .`"
	cd backend && go vet ./...

lint-frontend: ## Run frontend type checks
	@echo "$(GREEN)Linting frontend...$(NC)"
	cd frontend && npm run lint

build-backend: ## Build the backend binary
	@echo "$(GREEN)Building backend...$(NC)"
	cd backend && go build -o diputat-backend ./cmd/server

build-frontend: ## Build the frontend bundle
	@echo "$(GREEN)Building frontend...$(NC)"
	cd frontend && npm run build

seed: ## Describe the local seed source
	@echo "$(GREEN)Sample data lives in data/samples/*.json$(NC)"

migrate: ## Show the current migration sample
	@echo "$(GREEN)Current migration sample: data/migrations/001_initial_setup.json$(NC)"

db-shell: ## Connect to the MongoDB shell
	docker-compose exec mongodb mongosh -u admin -p password --authenticationDatabase admin diputat

status: ## Show Docker service status
	docker-compose ps
