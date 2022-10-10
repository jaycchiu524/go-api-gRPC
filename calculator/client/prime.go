package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jaycchiu524/go-api-grpc/calculator/proto"
)

func getPrimeNumbers(c pb.CalculatorServiceClient) {
	log.Println("get prime numbers was invoked")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := c.Prime(context.Background(), req); if err != nil {
		log.Fatalf("error while calling Prime RPC: %v", err)
	}

	count := 0

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Get prime %d: %v", count, msg.PrimeFactor)

		count += 1
	}
}