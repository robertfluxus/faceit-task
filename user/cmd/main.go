package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	userpb "github.com/robertfluxus/faceit-task/user/api"
	common "github.com/robertfluxus/faceit-task/user/common"
	userbusiness "github.com/robertfluxus/faceit-task/user/pkg/business"
	userdb "github.com/robertfluxus/faceit-task/user/pkg/db"
	usergrpc "github.com/robertfluxus/faceit-task/user/pkg/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type Opts struct {
	Port        string `short:"p" long:"port"`
	GatewayPort string `short:"g" long:"gateway_port"`
}

var opts Opts

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		log.Fatalf("failed to parse args: %w", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", opts.Port))
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
	go func() {
		log.Fatalln(grpcServer.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%s", opts.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to dial server: %w", err)
	}

	gatewayMux := runtime.NewServeMux()
	err = userpb.RegisterUserServiceHandler(context.Background(), gatewayMux, conn)
	if err != nil {
		log.Fatalf("Failed to register gateway: %w", err)
	}

	gatewayServer := &http.Server{
		Addr:    ":7001",
		Handler: gatewayMux,
	}
	log.Printf("Serving gateway on port %s", opts.GatewayPort)
	log.Fatalln(gatewayServer.ListenAndServe())
}
