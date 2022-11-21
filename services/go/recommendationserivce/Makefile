.PHONY: run
run: ## Run application with docker-compose
	@docker compose up -d --remove-orphans --build
	@docker compose logs -f app

.PHONY: down
down: ## Down docker containers from docker compose
	@docker compose down --remove-orphans

.PHONY: lint
lint: ## Lint golang code
	@docker run --rm -v ${PWD}/:/app -v ${PWD}/golangci.yml:/golangci.yml -w /app golangci/golangci-lint:latest golangci-lint run -v --fix --timeout=180s -c /golangci.yml

.PHONY: fieldalignment
fieldalignment: ## Detects and fixes structs that would use less memory if their fields were sorted (details: https://medium.com/@didi12468/golang-field-alignment-2e657e87668a)
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
	fieldalignment -test=false -fix=true $(PWD)/...

.PHONY: test-app
test-app:
	docker compose -f docker-compose.test.yaml up --build --exit-code-from test
	docker compose -f docker-compose.test.yaml rm -svf