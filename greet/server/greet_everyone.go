package main

import (
	"log"

	pb "github.com/jaycchiu524/go-api-grpc/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()
		if err != nil {
			break
		}

		firstName := req.FirstName
		result := "Hello " + firstName + "!"

		err = stream.Send(&pb.GreetResponse{
			Result: result,
		})

		if err != nil {
			return err
		}
	}

	return nil
}