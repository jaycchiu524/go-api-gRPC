package main

import (
	pb "github.com/jaycchiu524/go-api-grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	result := "Hello "

	for {
		req, err := stream.Recv()
		if err != nil {
			break
		}

		firstName := req.FirstName
		result += firstName + "! "
	}

	return stream.SendAndClose(&pb.GreetResponse{
		Result: result,
	})
}