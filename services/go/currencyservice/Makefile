.PHONY: run
run: ## Run application with docker-compose
	@docker compose up -d --remove-orphans --build
	@docker compose logs -f app

.PHONY: down
down: ## Down docker containers from docker compose
	@docker compose down --remove-orphans

.PHONY: lint
lint: ## Lint golang code
	@docker run --rm -v ${PWD}/:/app -v ${PWD}/golangci.yml:/golangci.yml -w /app golangci/golangci-lint:latest golangci-lint run -v --fix --timeout=90s -c /golangci.yml

.PHONY: test-app
test-app:
	docker compose -f docker-compose.test.yaml up --build --exit-code-from test
	docker compose -f docker-compose.test.yaml rm -svf
