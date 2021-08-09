package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/saskaradit/grpc-go-sandbox/client-streaming/calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello Im the client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer cc.Close()

	c := calcpb.NewCalculateServiceClient(cc)

	doClientStreaming(c)
}

func doClientStreaming(c calcpb.CalculateServiceClient) {
	fmt.Println("Starting to do a client streaming RPC...")

	requests := []*calcpb.CalculateLongRequest{
		&calcpb.CalculateLongRequest{
			Calculate: &calcpb.Calculate{
				Number: 20,
			},
		},
		&calcpb.CalculateLongRequest{
			Calculate: &calcpb.Calculate{
				Number: 12,
			},
		},
		&calcpb.CalculateLongRequest{
			Calculate: &calcpb.Calculate{
				Number: 15,
			},
		},
		&calcpb.CalculateLongRequest{
			Calculate: &calcpb.Calculate{
				Number: 45,
			},
		},
		&calcpb.CalculateLongRequest{
			Calculate: &calcpb.Calculate{
				Number: 600,
			},
		},
	}

	resStream, err := c.LongCalculate(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongCalculate: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending :%v\n", req)
		resStream.Send(req)
		time.Sleep(500 * time.Millisecond)
	}
	res, err := resStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving: %v", err)
	}
	fmt.Printf("LongCalculate Response, %v\n", res)
}
