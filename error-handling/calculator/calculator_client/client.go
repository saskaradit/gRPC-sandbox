package main

import (
	"context"
	"fmt"
	"log"

	"github.com/saskaradit/grpc-go-sandbox/error-handling/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hey Rad, I am the Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting SquareRoot RPC Unary..")
	// Correct call
	doErrorCall(c, 16)
	// Error Call
	doErrorCall(c, -2)

}

func doErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: n})
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			// Actual error from gRPC (user error)
			fmt.Println("Error message from server:", resErr.Message())
			fmt.Println(resErr.Code())
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("We propably sent a negative number")
				return
			}
		} else {
			log.Fatalln("Big Error calling squareroot", err)
			return
		}
	}
	fmt.Println("The square root of", n, "is", res.GetNumberRoot())
}
