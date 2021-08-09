package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/saskaradit/grpc-go-sandbox/unary/calculator/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculator(ctx context.Context, req *calcpb.CalculatorRequest) (*calcpb.CalculatorResponse, error) {
	fmt.Printf("Calculator function was invoked with %v\n", req)
	firstNum := req.GetCalculator().GetFirstNum()
	secondNum := req.GetCalculator().GetSecondNum()
	result := firstNum + secondNum
	res := &calcpb.CalculatorResponse{
		Result: result,
	}
	return res, nil

}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
