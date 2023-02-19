package main

import (
	"fmt"
	"log"

	pb "github.com/byegates/gRPC-go/greet/proto"
)

func (s *Server) GreetStream(in *pb.GreetRequest, stream pb.GreetService_GreetStreamServer) error {
	tag0 := tag + "[SS] "
	log.Printf("%v[Invoked] :{%v}\n", tag0, in)

	for i := 0; i < 5; i++ {
		stream.Send(&pb.GreetResponse{
			Result: fmt.Sprintf("Hello, %s, # %d 🃏", in.FirstName, i),
		})
	}

	log.Printf("\n\n")
	return nil
}
