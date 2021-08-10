package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/saskaradit/grpc-go-sandbox/deadlines/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello Im the client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnaryWithDeadline(c, 1*time.Second)
	doUnaryWithDeadline(c, 5*time.Second)
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Starting UnaryDeadline RPC...")
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Raditya",
			LastName:  "Saskara",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {

		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded")
			} else {
				fmt.Println("Unexpected error", statusErr)
			}
		} else {
			log.Fatalf("error while calling Greet RPC : %v", err)
		}
		return
	}
	log.Printf("Response from Greet: %v", res.Result)
}
