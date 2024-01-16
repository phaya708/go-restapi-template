seed:
	docker compose exec server go run cmd/db/seed/seed.go

reset:
	docker compose exec server go run cmd/db/reset/reset.go

test-all:
	docker compose exec server go test -v ./...

test:
	docker compose exec server go test -v $(path)

bash:
	docker compose exec server bash