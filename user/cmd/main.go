package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	userpb "github.com/robertfluxus/faceit-task/user/api"
	common "github.com/robertfluxus/faceit-task/user/common"
	userbusiness "github.com/robertfluxus/faceit-task/user/pkg/business"
	userdb "github.com/robertfluxus/faceit-task/user/pkg/db"
	usergrpc "github.com/robertfluxus/faceit-task/user/pkg/grpc"

	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type Opts struct {
	Port int `short:"p" long:"port"`
}

var opts Opts

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		log.Fatalf("failed to parse args: %w", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.Port))
	if err != nil {
		log.Fatal("failed to listen: %w", err)
	}

	connURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
	)

	db, err := common.ConnectPostgress(context.Background(), common.NewDefaultConnectionOptions(connURL))
	if err != nil {
		log.Fatalf("Failed connecting to the database: %w", err)
	}
	defer db.Close()
	userRespository := userdb.New(db)

	rabbitConnURL := "amqp://rabbitmq:rabbitmq@rabbithost:5672/"
	rabbitConn, err := common.ConnectRabbitMQ(context.Background(), common.NewDefaultConnectionOptions(rabbitConnURL))
	if err != nil {
		log.Printf("Failed to connect to rabbit mq: %w", err)
	}
	rabbit, err := userbusiness.NewRabbitMQ(rabbitConn)
	if err != nil {
		log.Printf("Failed to create channel: %w", err)
	}
	rabbit.CreateQueue("USER_UPDATES")
	userService := userbusiness.NewUserService(userRespository, db, rabbit)

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, usergrpc.NewUserServiceHandler(userService))

	log.Printf("Initializing gRPC server on port %d", opts.Port)
	grpcServer.Serve(lis)

}
