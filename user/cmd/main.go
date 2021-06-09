package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	userpb "github.com/robertfluxus/faceit-task/user/api"
	userbusiness "github.com/robertfluxus/faceit-task/user/pkg/business"
	userdb "github.com/robertfluxus/faceit-task/user/pkg/db"
	usergrpc "github.com/robertfluxus/faceit-task/user/pkg/grpc"

	"github.com/jessevdk/go-flags"
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

	connURL := "test"
	db, err := sql.Open("postgres", connURL)
	userRespository := userdb.New(db)

	userService := userbusiness.NewUserService(userRespository, db)

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, usergrpc.NewUserService(userService))

	log.Printf("Initializing gRPC server on port %d", opts.Port)
	grpcServer.Serve(lis)

}
