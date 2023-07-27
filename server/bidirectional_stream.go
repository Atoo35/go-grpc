package main

import (
	"io"
	"log"

	pb "github.com/Atoo35/go-grpc/proto"
)

func (s *server) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received: %v", req.GetName())
		stream.Send(&pb.HelloResponse{
			Message: "Hello " + req.GetName(),
		})
	}
}
