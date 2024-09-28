# define variables
DB_NAME=fiber-jwt-example
DB_USER=postgres
DB_PASSWORD=root
DB_PORT=5432


init-db:
	@echo "Initializing database $(DB_NAME)..."
	@if [ $$(docker ps -a -q -f name=$(DB_NAME)) ]; then \
		if [ $$(docker ps -q -f name=$(DB_NAME)) ]; then \
			echo "Database $(DB_NAME) is already running."; \
		else \
			echo "Starting existing database container $(DB_NAME)..."; \
			docker start $(DB_NAME); \
		fi \
	else \
		echo "Creating and starting new database container $(DB_NAME)..."; \
		docker run --name $(DB_NAME) -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -e POSTGRES_DB=$(DB_NAME) -p $(DB_PORT):5432 -d postgres; \
	fi