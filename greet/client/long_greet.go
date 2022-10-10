package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jaycchiu524/go-api-grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Jay"},
		{FirstName: "Wanying"},
		{FirstName: "Marcus"},
		{FirstName: "Blue"},
		{FirstName: "Ivan"},
	}
	stream, err := c.LongGreet(context.Background())

	if(err != nil) {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}

	log.Printf("LongGreet Response: %v", res)
}