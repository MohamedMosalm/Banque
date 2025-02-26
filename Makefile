postgres:
	docker run --name postgres-container -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres

createdb:
	docker exec -it postgres-container createdb --username=postgres --owner=postgres banque

dropdb:
	docker exec -it postgres-container dropdb banque

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/banque?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/banque?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	@go test -v -cover ./...

server:
	@go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/MohamedMosalm/banque/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock