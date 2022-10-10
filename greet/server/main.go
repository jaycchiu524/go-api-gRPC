package main

import (
	"log"
	"net"

	pb "github.com/jaycchiu524/go-api-grpc/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if(err != nil) {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()

	// Register the service
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}