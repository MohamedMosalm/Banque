DB_URL=postgresql://postgres:password@localhost:5432/banque?sslmode=disable

postgres:
	docker run --name postgres-container -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres

createdb:
	docker exec -it postgres-container createdb --username=postgres --owner=postgres banque

dropdb:
	docker exec -it postgres-container dropdb banque

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

db_docs:
	npx dbdocs build doc/db.dbml

db_schema:
	npx dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	@go test -v -cover ./...

server:
	@go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/MohamedMosalm/banque/db/sqlc Store

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: postgres createdb dropdb migrateup migratedown db_docs db_schema sqlc server mock proto