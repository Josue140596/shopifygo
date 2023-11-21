createdb:
	docker exec -it postgresShopify createdb --username=root --owner=root shopify_db
docker-up:
	docker-compose up -d
docker-down:
	docker-compose down
docker-start:
	docker-compose start
docker-restart:
	docker-compose restart
docker-stop:
	docker-compose stop
db-shell:
	docker-compose exec db psql -U root shopify_db
dropdb:
	docker exec -it postgresShopify dropdb shopify_db
migrateup:
	migrate -path internal/db/general_migration -database "postgres://root:secret@localhost:5432/shopify_db?sslmode=disable" -verbose up
migratedown:
	migrate -path internal/db/general_migration -database "postgres://root:secret@localhost:5432/shopify_db?sslmode=disable" -verbose down
test: 
	go test -v -cover ./...
serve: 
	go run ./cmd/user-service/main.go

.PHONY: postgres createdb dropdb migrateup serve test docker-up docker-down docker-start docker-restart docker-stop db-shell containerDB migratedown