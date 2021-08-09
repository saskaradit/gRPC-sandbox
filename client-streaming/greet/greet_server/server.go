package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/saskaradit/grpc-go-sandbox/client-streaming/greet/greetpb"
	grpc "google.golang.org/grpc"
)

type server struct{}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("LongGreet function was invoked with a streaming request")
	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// The client streaming has ended
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading the client stream %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += firstName + " "
	}
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
