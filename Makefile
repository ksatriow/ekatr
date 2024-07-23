SERVICE_NAME = ekatr

.PHONY: help
help:
	@echo "Usage:"
	@echo "  make build       Build the Docker image"
	@echo "  make up          Start the Docker containers"
	@echo "  make down        Stop and remove the Docker containers"
	@echo "  make logs        View the logs from the Docker containers"
	@echo "  make ps          List the Docker containers"
	@echo "  make clean       Remove Docker containers and volumes"

.PHONY: build
build:
	docker-compose build

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: logs
logs:
	docker-compose logs -f $(SERVICE_NAME)

.PHONY: ps
ps:
	docker-compose ps

.PHONY: clean
clean:
	docker-compose down -v
