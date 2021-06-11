# Faceit Take Home Challenge

## How to run

This project can be deployed locally using docker compose

1. `make buildfaceit`
2. `make compose`

To remove:  `make removecompose`

## How to interact

1. Using cURL/Postman.
The service exposes a REST API. Example command: 

    `curl -X GET http://127.0.0.1:7001/v1/user/testuser`

2. Using gRPCox.
To start a gRPCox container run `make grpcox`.
On your browser go to `localhost:6969`.
On gRPC Server Target entert `localhost:7000`.
Select use local proto (the service doesn't not use reflection).
Select the user service from the dropdown and the desired method.

## How to view the data

1. To view the postgres database, you can use TablePlus or using psql:

    `docker exec -it postgres12 psql -U root faceit`

2. To view the RabbitMQ dashboard:
On your browser go to: `http://localhost:15672/`.
Go to `Queues`.
Select `USER_UPDATES`.
Click `Get Message(s)`

