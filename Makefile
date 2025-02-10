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

