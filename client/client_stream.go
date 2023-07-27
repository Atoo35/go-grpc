package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Atoo35/go-grpc/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for _, name := range names.Names {
		if err := stream.Send(&pb.HelloRequest{
			Name: name,
		}); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		log.Printf("Greeting sent: %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}
	log.Printf("%v", res.GetMessages())
}
