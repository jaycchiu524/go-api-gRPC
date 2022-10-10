package main

import (
	"context"
	"log"

	pb "github.com/jaycchiu524/go-api-grpc/greet/proto"
)
	
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	// in is the request
	log.Printf("Received: %v", in.FirstName)
	return &pb.GreetResponse {
		Result: "Hello " + in.FirstName,
	}, nil
}