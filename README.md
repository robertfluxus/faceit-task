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

## What can be improved:
1. More tests, due to the timeframe I have only added tests for the grpc layer. These contain mocks as well. The same approach can be taken for all of the layers. A test suite struct can be created as well with functionality to set up tests and tear down tests.
2. The password can be hashed and stored as a hash instead of it's actual form.
3. Transport layer security can be added for the grpc, credentials can be generated etc.
4. The RabbitMQ common library can be extended to include better producers and consumers as well. Potentially for bigger services that have a high throughput kafka can be used.
5. Kubernetes can be used instead of docker compose, this would allow for horizontal scalability (maybe using HPA), services can be created, additional security layers  can be added, more observability.
6. Better logging can be implemented, maybe structured logging using `https://github.com/sirupsen/logrus`.
7. Idempotency can be implemented better using request ids.  