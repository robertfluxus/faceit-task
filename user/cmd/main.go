package main

import (
	"fmt"
	"log"
	"net"
	"os"

	userpb "github.com/robertfluxus/faceit-task/user/api"

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

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer)

	log.Printf("Initializing gRPC server on port %d", opts.Port)
	grpcServer.Serve(lis)

}
