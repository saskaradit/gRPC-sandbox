package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/saskaradit/grpc-go-sandbox/server-streaming/greet/greetpb"
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

	doServerStreaming(c)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Server Streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rad",
			LastName:  "Saskara",
		},
	}
	reStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetMany Times: %v", err)
	}
	for {
		msg, err := reStream.Recv()
		if err == io.EOF {
			// reached the end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())

	}

}
