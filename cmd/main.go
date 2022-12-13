package main

import (
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
	l, err := net.Listen("tcp", hostport)
	if err != nil {
		log.Fatalf("Error starting TCP listener: %v", err)
	}
	// create, register, and start a gRPC server
	s := grpc.NewServer()
	grpctest.RegisterExampleServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Error starting to serve gRPC: %v", err)
	}
}