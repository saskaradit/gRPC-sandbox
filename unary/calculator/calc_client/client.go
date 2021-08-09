package main

import (
	"context"
	"fmt"
	"log"

	"github.com/saskaradit/grpc-go-sandbox/unary/calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hey Rad, I am the Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer cc.Close()
	c := calcpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting Unary..")
	req := &calcpb.CalculatorRequest{
		Calculator: &calcpb.Calculator{
			FirstNum:  20,
			SecondNum: 100,
		},
	}

	res, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling calculator RPC: %v", err)
	}
	log.Printf("Response from Calculator: %v", res.Result)
}
