package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/byegates/gRPC-go/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	tag0 := tag + "[Bi] "
	log.Printf("%v[Invoked]\n", tag0)

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("%v[Error] while creating stream: %v\n", tag0, err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("%v[Send] {%v}\n", tag0, req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
		log.Printf("%v[Send] [end] to stream\n\n", tag0)
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("%v[Error] [Recv] from stream: %v\n", tag0, err)
				break
			}

			log.Printf("%v[Recv] %s\n", tag0, res.Result)
		}

		close(waitc)
		// log.Printf("%vReceiving stream End\n\n", tag0)
	}()

	<-waitc
}
