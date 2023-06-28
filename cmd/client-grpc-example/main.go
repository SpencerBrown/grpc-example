package main

import (
	"context"
	"log"
	"time"

	"github.com/SpencerBrown/grpc-example/grpctest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const hostport = "localhost:9000"

func main() {
	// Connect to gRPC server with no credentials.
	connection, err := grpc.Dial(hostport, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to gRPC server: %v", err)
	}
	// Create gRPC client instance 
	client := grpctest.NewExampleClient(connection)
	// Send request to server with a one second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	response, err := client.Ping(ctx, &grpctest.PingRequest{})
	if err != nil {
		cancel()
		connection.Close()
		log.Fatalf("Error sending request to server and getting response: %v", err)
	}
	log.Printf("Response: %s", response.GetStatus())
}
