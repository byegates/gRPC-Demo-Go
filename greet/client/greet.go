package main

import (
	"context"
	"log"

	pb "github.com/byegates/gRPC-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	tag0 := tag + "[Greet] "
	log.Printf("%v[Invoked]\n", tag0)
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Qiuchen",
	})

	if err != nil {
		log.Fatalf("%vCould not greet: %v\n", tag0, err)
	}

	log.Printf("%v[Recv]: %s\n\n", tag0, res.Result)
}
