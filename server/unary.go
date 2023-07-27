package main

import (
	"context"

	pb "github.com/Atoo35/go-grpc/proto"
)

func (s *server) SayHello(ctx context.Context, in *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello from the server!",
	}, nil
}
