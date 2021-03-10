docker-build:
	@docker build -t frete-analyzer-api .

start-deps:
	@echo "Starting dependencies..."
	docker-compose up -d
	@until docker exec postgres-fr pg_isready; do echo 'Waiting Postgres...' && sleep 5; done
	@echo "Dependencies started successfully."

stop-deps:
	@echo "Stopping dependencies..."
	docker-compose down
	@echo "Dependencies stopped successfully."

restart-deps:
	make stop-deps
	make start-deps

run-server:
	go run main.go server

