DB_URL=postgres://app:app@localhost:5432/colab?sslmode=disable
MIGRATIONS=./migrations

.PHONY: db-up db-down test run migrate-up migrate-down migrate-version

db-up:
	docker compose up -d

db-down:
	docker compose down

run:
	go run ./cmd/api

test:
	go test ./...

migrate-up:
	migrate -path $(MIGRATIONS) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS) -database "$(DB_URL)" down 1

migrate-version:
	migrate -path $(MIGRATIONS) -database "$(DB_URL)" version
