package main

import (
	"context"
	"log"
	"time"

	pb "github.com/byegates/gRPC-go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	tag0 := tag + "[CS] "
	log.Printf("%v[Invoked]\n", tag0)

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("%v[Error] creating stream: %v\n", tag0, err)
	}

	for _, req := range reqs {
		log.Printf("%v[Sent] to stream: {%v}\n", tag0, req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v[Error] [Recv]: %v\n", tag0, err)
	}

	log.Printf("\n%v[Recv]:\n\n%s\n", tag0, res.Result)
}
