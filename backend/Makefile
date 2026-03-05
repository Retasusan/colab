.PHONY: db-up db-down test run

db-up:
	docker compose up -d

db-down:
	docker compose down

run:
	go run ./cmd/api

test:
	go test ./...
