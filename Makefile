.PHONY: up
up:
	@echo Starting Docker images...
	docker-compose up -d
	@echo Docker images started!

.PHONY: up/build
up/build:
	@echo Starting Docker images...
	docker-compose up --build -d
	@echo Docker images started!

.PHONY: down
down:
	@echo Stopping Docker images...
	docker-compose down
	@echo Docker images removed!

## migrate/new service=$1 name=$2: create a new database migration
.PHONY: migrate/new
migrate/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext .sql -dir ./${service}/migrations ${name}

