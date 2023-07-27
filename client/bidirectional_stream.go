package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Atoo35/go-grpc/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming: %v", err)
			}
			log.Printf("Greeting Received: %s", message.Message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		if err := stream.Send(&pb.HelloRequest{
			Name: name,
		}); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		log.Printf("Greeting sent: %s", name)
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional streaming finished")
}
