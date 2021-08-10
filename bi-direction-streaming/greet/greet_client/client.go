package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/saskaradit/grpc-go-sandbox/bi-direction-streaming/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello Im the client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doBiDiStreaming(c)
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a BiDi streaming RPC...")

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Rad",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Saskara",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Bread",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Zeref",
			},
		},
	}

	// Create stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalln("Error while creating stream", err)
		return
	}

	waitCh := make(chan struct{})
	// send bunch of messages to the client
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Println("Sending message", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// Receive a bunch of message from the client
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("Error while receiving", err)
				break
			}
			fmt.Println("Received", res.GetResult())
		}
		close(waitCh)
	}()

	// Block until done
	<-waitCh
}
