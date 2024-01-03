postgres:
	docker run --name hypergofram-postgres-1 -e POSTGRES_PASSWORD=secret123 -p 5432:5432 -d postgres:16.1-alpine

createdb: 
	docker exec -it hypergofram-postgres-1 createdb --username=root --owner=root hypergofram

dropdb:
	docker exec -it hypergofram-postgres-1 dropdb hypergofram

migrateup:
	migrate -path internal/db/migrations -database "postgresql://root:secret123@0.0.0.0:5432/hypergofram?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migrations -database "postgresql://root:secret123@0.0.0.0:5432/hypergofram?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown