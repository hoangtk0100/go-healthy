include app.env
export

network:
	docker network create $(NETWORK_NAME)

postgres:
	docker run --name $(DB_CONTAINER_NAME) --network $(NETWORK_NAME) -p $(DB_PORT):5432 -e POSTGRES_USER=$(DB_USERNAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d $(PSQL_IMAGE)

createdb:
	docker exec -it $(DB_CONTAINER_NAME) createdb --username=$(DB_USERNAME) --owner=$(DB_USERNAME) $(DB_NAME)

dropdb:
	docker exec -it $(DB_CONTAINER_NAME) dropdb $(DB_NAME)

db:
	docker exec -it $(DB_CONTAINER_NAME) psql -U root $(DB_NAME)

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down 1

sqlc:
	sqlc generate

.PHONY: network postgres createdb dropdb db new_migration migrateup migratedown migrateup1 migratedown1 sqlc