package main

import (
	"fmt"
	"log"

	pb "github.com/jaycchiu524/go-api-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes was invoked: %v", in)

	for i := 0; i<10; i++ {
		res :=fmt.Sprintf("Hello %v, number %v", in.FirstName, i);

		err := stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			return err 
		}
	}

	return nil
}