package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/saskaradit/grpc-go-sandbox/bi-direction-streaming/maximum/maxpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) MaximumNumber(stream maxpb.MaximumService_MaximumNumberServer) error {
	fmt.Println("MaximumNumber function was invoked with a streaming request")
	maxList := []int32{}
	var max int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Stream ended
			return nil
		}
		if err != nil {
			log.Fatalln("Error while reading client stream", err)
		}
		num := req.GetNumber()
		if num > max {
			max = num
			maxList = append(maxList, max)
		}
		err = stream.Send(&maxpb.MaximumNumberResponse{
			Result: maxList,
		})
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}
	fmt.Println("Connected to server")

	s := grpc.NewServer()
	maxpb.RegisterMaximumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve", err)
	}
}
