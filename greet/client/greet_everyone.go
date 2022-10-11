package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/jaycchiu524/go-api-grpc/greet/proto"
)

func doGreetEveryone (c pb.GreetServiceClient) {
	log.Println("do greet everyone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while calling GreetEveryone RPC: %v", err)
	}

	// waitc := make(chan struct{})

	// go func() {
	// 	for {
	// 		res, err := stream.Recv()
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		if err != nil {
	// 			log.Fatalf("error while reading stream: %v", err)
	// 			break
	// 		}
	// 		log.Printf("Response from GreetEveryone: %v", res.Result)
	// 	}
	// 	close(waitc)
	// }()

	// names := []string{"Jay", "Wanying", "Ivan", "Marcus", "Blue"}
	// for _, name := range names {
	// 	log.Printf("Sending name: %v", name)
	// 	stream.Send(&pb.GreetRequest{
	// 		FirstName: name,
	// 	})
	// 	time.Sleep(1000 * time.Millisecond)
	// }

	// stream.CloseSend()
	// <-waitc

	waitc := make(chan struct{})

	names := []string{"Jay", "Wanying", "Ivan", "Marcus", "Blue"}

	// send a bunch of messages to the client (go routine)
	go func() {
		for _, name := range names {
			log.Printf("Sending name: %v", name)
			stream.Send(&pb.GreetRequest{
				FirstName: name,
			})

			// delay to simulate a streaming request
			time.Sleep(1000 * time.Millisecond)
		}
		// Close stream after sending all names
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if(err != nil) {
				log.Fatalf("error while receiving: %v", err)
			}

			log.Printf("Response from GreetEveryone: %v", res.Result)
		}
		close(waitc)
	}()
	
	// block until everything is done
	<-waitc
}