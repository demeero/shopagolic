MONGO_URL ?= mongodb://localhost:27017/shopagolic-catalog

.PHONY: run
run: ## Run application with docker-compose
	@docker compose up -d --remove-orphans --build
	@docker compose logs -f app

.PHONY: run-env
run-env: ## Run app env: mongo, mongo-express, etc to run the app locally
	@docker compose -f docker-compose.yaml up -d mongo mongo-express

.PHONY: down
down: ## Down docker containers from docker compose
	@docker compose down --remove-orphans

.PHONY: lint
lint: ## Lint golang code
	@docker run --rm -v ${PWD}/:/app -v ${PWD}/golangci.yml:/golangci.yml -w /app golangci/golangci-lint:latest golangci-lint run -v --fix --timeout=90s -c /golangci.yml

.PHONY: test-app
test-app:
	docker compose -f docker-compose.test.yaml up -d --build --exit-code-from test
	docker compose -f docker-compose.test.yaml rm -svf

.PHONY: migration
migration: ## Create a new up and down migration
	@docker run -v ${PWD}/migrations:/migration:rw --network host migrate/migrate \
		create -ext json -dir /migration $(name)

.PHONY: migration-mongo-up
migration-mongo-up: scale_up=
migration-mongo-up: ## Apply all or N up migrations to MongoDB
	@docker run --rm -v ${PWD}/migrations:/migration:ro --network host migrate/migrate \
  		-source file://migration -database ${MONGO_URL} up ${scale_up}

.PHONY: migration-mongo-down
migration-mongo-down: scale_down=1
migration-mongo-down: ## Apply 1 or N down migrations to MongoDB
	@docker run --rm -v ${PWD}/migrations:/migration:ro --network host migrate/migrate \
		-source file://migration -database ${MONGO_URL} down ${scale_down}