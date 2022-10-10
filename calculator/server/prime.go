package main

import (
	"log"

	pb "github.com/jaycchiu524/go-api-grpc/calculator/proto"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Printf("Received Generating Prime number for: %v", in.Number)

	var k int32 = 2
	N := in.Number
	c := 0

	for N > 1 {
		if N%k == 0 {

			err := stream.Send(&pb.PrimeResponse{
				PrimeFactor: k,
			})

			if err != nil {
				return err
			}

			N = N / k
			c += 1
		} else {
			k = k + 1
		}
	}

	return nil
}