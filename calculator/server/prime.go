package main

import (
	"log"

	pb "github.com/jaycchiu524/go-api-grpc/calculator/proto"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Printf("Received Generating Prime number for: %v", in.Number)

	divisor := int32(2)
	number := in.Number

	for number > 1 {
		if number%divisor == 0 {

			err := stream.Send(&pb.PrimeResponse{
				PrimeFactor: divisor,
			})

			if err != nil {
				return err
			}

			number /= divisor
		} else {
			divisor +=  1
		}
	}

	return nil
}