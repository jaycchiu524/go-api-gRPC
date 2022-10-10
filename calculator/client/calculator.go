package main

import (
	"context"
	"log"

	pb "github.com/jaycchiu524/go-api-grpc/calculator/proto"
)

func doSum(c pb.SumServiceClient) {
	log.Println("do sum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 1,
		B: 5,
	}); if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Printf("Sum: %d\n", res.Sum)
}