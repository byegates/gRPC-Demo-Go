package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/byegates/gRPC-go/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	tag0 := tag + "[BI] "
	log.Printf("%v[Invoked]\n", tag0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Printf("%v[End]\n\n", tag0)
			return nil
		}

		if err != nil {
			log.Fatalf("%v[Error] [Recv] from stream: %v\n", tag0, err)
		}

		err = stream.Send(&pb.GreetResponse{Result: fmt.Sprintf("Hello, %s 🃏", req.FirstName)})

		if err != nil {
			log.Fatalf("%v[Error] [Send] to stream: %v\n", tag0, err)
		}
	}
}
