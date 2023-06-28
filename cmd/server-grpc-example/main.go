package main

import (
	"context"
	"log"
	"net"

	"github.com/SpencerBrown/grpc-example/grpctest"
	"google.golang.org/grpc"
)

const hostport = "localhost:9000"

// server is used to implement grpctest.Example interface.
// All implementations must embed UnimplementedExampleServer for forward compatibility.
type server struct {
	grpctest.UnimplementedExampleServer
}

func main() {
	// set up grpc server. first, listen on the port using net package
	listener, err := net.Listen("tcp", hostport)
	if err != nil {
		log.Fatalf("Error starting TCP listener: %v", err)
	}
	// create, register, and start a gRPC server
	myServer := grpc.NewServer()
	grpctest.RegisterExampleServer(myServer, &server{})
	log.Printf("Starting server on %s; use Ctrl+C to stop it", hostport)
	if err := myServer.Serve(listener); err != nil {
		log.Fatalf("Error serving gRPC: %v", err)
	}
}

func (myServer *server) Ping(ctx context.Context, request *grpctest.PingRequest) (*grpctest.PingStatus, error) {
	log.Printf("Ping request received: %v", request.String())
	return &grpctest.PingStatus{
		Status: "All good here",
	}, nil
}
