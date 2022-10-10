package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jaycchiu524/go-api-grpc/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("do greet many times was invoked")

	req := &pb.GreetRequest{
		FirstName: "Jay",
	}

	stream, err := c.GreetManyTimes(context.Background(), req); if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	// for without a condition will loop repeatedly until you break out of the loop 
	// or return from the enclosing function

	for {
		// receive a message from the server
		msg, err := stream.Recv()

		// io.EOF is the error returned by Read when no more input is available.
		if err == io.EOF {
			break
		}

		// if err is not nil and err is not io.EOF
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Response from GreetManyTimes: %v", msg.Result)
	}
}