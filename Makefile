.PHONY: help build up down logs clean restart status

# Default target
help:
	@echo "Queue Processing Service - Docker Commands"
	@echo "========================================="
	@echo ""
	@echo "Commands:"
	@echo "  up        - Start all services (PostgreSQL, Redis, Go app)"
	@echo "  down      - Stop all services"
	@echo "  build     - Build the Go application Docker image"
	@echo "  logs      - Show logs from all services"
	@echo "  status    - Show status of all services"
	@echo "  restart   - Restart all services"
	@echo "  clean     - Stop services and remove volumes (WARNING: data loss)"
	@echo "  help      - Show this help message"

# Start all services
up:
	docker-compose up -d

# Stop all services
down:
	docker-compose down

# Build the application
build:
	docker-compose build

# Show logs
logs:
	docker-compose logs -f

# Show service status
status:
	docker-compose ps

# Restart all services
restart:
	docker-compose restart

# Clean everything (including data)
clean:
	docker-compose down -v
	docker system prune -f
