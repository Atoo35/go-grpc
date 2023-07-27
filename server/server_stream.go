package main

import (
	"log"
	"time"

	pb "github.com/Atoo35/go-grpc/proto"
)

func (s *server) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Received: %v", req.GetNames())
	for _, name := range req.Names {
		if err := stream.Send(&pb.HelloResponse{
			Message: "Hello " + name,
		}); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
