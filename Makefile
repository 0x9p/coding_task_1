.PHONY: test

POSTGRES_DB_USERNAME=developer
POSTGRES_DB_PASSWORD=qwaszx
POSTGRES_DB_NAME=coding_task_1

run-svc:
	go run ./cmd/http/main.go

run-pg:
	docker-compose -f docker-compose.yml up -d postgres

run-pg-migrations:
	migrate -path ./migrations/postgres -database "postgres://${POSTGRES_DB_USERNAME}:${POSTGRES_DB_PASSWORD}@localhost/${POSTGRES_DB_NAME}?sslmode=disable" up

down-pg-migrations:
	migrate -path ./migrations/postgres -database "postgres://${POSTGRES_DB_USERNAME}:${POSTGRES_DB_PASSWORD}@localhost/${POSTGRES_DB_NAME}?sslmode=disable" down

down:
	docker-compose down

test:
	go test -v -cover ./internal/...

deps:
	go mod download

clean-deps:
	go mod tidy

.DEFAULT_GOAL := test
