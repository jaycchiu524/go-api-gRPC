package main

import (
	"context"
	"log"

	pb "github.com/jaycchiu524/go-api-grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("do greet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Jay",
	})

	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}