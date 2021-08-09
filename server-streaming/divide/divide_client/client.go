package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/saskaradit/grpc-go-sandbox/server-streaming/divide/dividepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello im the client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer cc.Close()

	c := dividepb.NewDivideServiceClient(cc)

	doServerStreaming(c)
}

func doServerStreaming(c dividepb.DivideServiceClient) {
	fmt.Println("Starting server streaming RPC....")
	req := &dividepb.DivideManyTimesRequest{
		Divide: &dividepb.Divide{
			Number: 120,
		},
	}
	resStream, err := c.DivideManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error whiule calling DivideManyTimes: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// Reached the end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream: %v", err)
		}
		log.Printf("Repsonse from DivideManyTimes: %v", msg.GetResult())
	}
}
