.PHONY: help build run test docker-up docker-down docker-logs clean

help:
	@echo "Weather App - Available Commands"
	@echo ""
	@echo "Development:"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application locally"
	@echo "  make test         - Run unit tests"
	@echo "  make test-verbose - Run unit tests with verbose output"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-up    - Start containers with docker-compose"
	@echo "  make docker-down  - Stop containers"
	@echo "  make docker-logs  - View container logs"
	@echo ""
	@echo "Deployment:"
	@echo "  make deploy       - Deploy to Google Cloud Run"
	@echo ""
	@echo "Cleanup:"
	@echo "  make clean        - Remove build artifacts"

build:
	@echo "Building application..."
	go build -o weather-app ./cmd/weatherapp

run: build
	@echo "Running application..."
	./weather-app

test:
	@echo "Running tests..."
	go test -v ./...

test-verbose:
	@echo "Running tests with coverage..."
	go test -v -cover ./...

docker-build:
	@echo "Building Docker image..."
	docker-compose build

docker-up: docker-build
	@echo "Starting containers..."
	docker-compose up -d
	@echo "Application available at http://localhost:8080"

docker-down:
	@echo "Stopping containers..."
	docker-compose down

docker-logs:
	@echo "Showing container logs..."
	docker-compose logs -f app

deploy:
	@echo "Deploying to Google Cloud Run..."
	gcloud run deploy weather-app \
		--source . \
		--platform managed \
		--region us-central1 \
		--allow-unauthenticated \
		--set-env-vars WEATHER_API_KEY=$(WEATHER_API_KEY)

clean:
	@echo "Cleaning up..."
	rm -f weather-app
	go clean
	docker-compose down -v

.DEFAULT_GOAL := help
