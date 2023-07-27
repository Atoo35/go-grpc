package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Atoo35/go-grpc/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", res.Message)
	}
	log.Printf("Streaming finished")
}
