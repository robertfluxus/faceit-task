.ONESHELL:

protogen:
	cd user/api
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		user.proto

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root faceit

dropdb:
	docker exec -it postgres12 dropdb faceit

migrateup:
	migrate -path user/pkg/db/migration -database "postgresql://root:secret@localhost:5432/faceit?sslmode=disable" -verbose up

migratedown:
	migrate -path user/pkg/db/migration -database "postgresql://root:secret@localhost:5432/faceit?sslmode=disable" -verbose down

removepostgres:
	docker stop postgres12
	docker rm postgres12

.PHONY: protogen