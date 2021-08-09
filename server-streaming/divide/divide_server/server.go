package main

import (
	"fmt"
	"log"
	"net"

	"github.com/saskaradit/grpc-go-sandbox/server-streaming/divide/dividepb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) DivideManyTimes(req *dividepb.DivideManyTimesRequest, stream dividepb.DivideService_DivideManyTimesServer) error {
	fmt.Printf("DivideManyTimes was invoked with: %v\n", req)
	num := req.GetDivide().GetNumber()
	var div int32 = 2
	for num > 1 {
		if num%div == 0 {
			stream.Send(&dividepb.DivideManyTimesResponse{
				Result: div,
			})
			num = num / div
		} else {
			div = div + 1
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	dividepb.RegisterDivideServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
