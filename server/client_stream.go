package main

import (
	"io"
	"log"

	pb "github.com/Atoo35/go-grpc/proto"
)

func (s *server) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Received: %v", req.GetName())
		messages = append(messages, "Hello", req.GetName())
	}
}
