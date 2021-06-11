.ONESHELL:

protogen:
	cd user/api
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=. --grpc-gateway_opt paths=source_relative \
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

buildfaceit:
	cd user/cmd
	CGO_ENABLED=0 GOOS=linux go build -o ../k8s/faceittask

compose:
	cd user/k8s
	docker-compose up --build --force-recreate -d
	docker-compose logs -f -t

removecompose:
	docker-compose down

rabbit:
	docker run -it -d --hostname rabbit --name rabbit --rm -ti --net="host" rabbitmq:3.8-management 

removerabbit:
	docker stop rabbit

grpcox:
	docker run --net=host -p 6969:6969 -v $(pwd)/logs/log:/log -d gusaul/grpcox 

mockuserservice:
	mockgen -destination user/pkg/grpc/mocks/user_service_mock.go \
		github.com/robertfluxus/faceit-task/user/pkg/grpc UserService
