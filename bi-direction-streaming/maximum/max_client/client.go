package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/saskaradit/grpc-go-sandbox/bi-direction-streaming/maximum/maxpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello im the client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to connect", err)
	}

	defer cc.Close()

	c := maxpb.NewMaximumServiceClient(cc)

	doBiDiStreaming(c)
}

func doBiDiStreaming(c maxpb.MaximumServiceClient) {
	fmt.Println("Started a bi directional streaming RPC...")

	requests := []*maxpb.MaximumNumberRequest{
		&maxpb.MaximumNumberRequest{
			Number: 1,
		},
		&maxpb.MaximumNumberRequest{
			Number: 5,
		},
		&maxpb.MaximumNumberRequest{
			Number: 3,
		},
		&maxpb.MaximumNumberRequest{
			Number: 6,
		},
		&maxpb.MaximumNumberRequest{
			Number: 2,
		},
		&maxpb.MaximumNumberRequest{
			Number: 20,
		},
	}

	// Create stream
	stream, err := c.MaximumNumber(context.Background())
	if err != nil {
		log.Fatalln("Error while creating stream", err)
		return
	}

	waitCh := make(chan struct{})
	// send messages to client
	go func() {
		for _, req := range requests {
			fmt.Println("Sending message", req)
			stream.Send(req)
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// Receive messages
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("Error while receiving", err)
				break
			}
			fmt.Println("Received", res.GetResult())
		}
		close(waitCh)
	}()

	// Block until done
	<-waitCh
}
