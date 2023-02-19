package main

import (
	"context"
	"io"
	"log"

	pb "github.com/byegates/gRPC-go/greet/proto"
)

func doGreetStream(c pb.GreetServiceClient) {
	tag0 := tag + "[SS]"
	log.Printf("%v[Invoked]\n", tag0)
	stream, err := c.GreetStream(context.Background(), &pb.GreetRequest{
		FirstName: "Qiuchen",
	})

	if err != nil {
		log.Fatalf("%v[Error] while creating stream: %v\n", tag0, err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("%v[Error] [Recv] from stream: %v\n", tag0, err)
		}

		log.Printf("%v[Recv] %s\n", tag0, msg.Result)
	}

	log.Printf("%v[Recv] stream End\n\n", tag0)
}
