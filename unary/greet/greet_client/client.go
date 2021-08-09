package main

import (
	"context"
	"fmt"
	"log"

	"github.com/saskaradit/grpc-go-sandbox/unary/greet/greetpb"
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
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Unary...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Raditya",
			LastName:  "Saskara",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC : %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
