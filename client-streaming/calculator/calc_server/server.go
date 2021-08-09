package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/saskaradit/grpc-go-sandbox/client-streaming/calculator/calcpb"
	grpc "google.golang.org/grpc"
)

type server struct{}

func (*server) LongCalculate(stream calcpb.CalculateService_LongCalculateServer) error {
	fmt.Println("LongCalculate function was invoked with a streaming request")
	var result int32 = 0
	var count int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// The client streaming has ended
			return stream.SendAndClose(&calcpb.CalculateLongResponse{
				Result: result / count,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading the client %v", err)
		}
		count++
		result += req.GetCalculate().GetNumber()
	}

}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Connected to server")
	s := grpc.NewServer()
	calcpb.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
