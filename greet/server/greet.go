package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/byegates/gRPC-go/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("%v[Greet] [Invoked] : {%v}\n\n", tag, in)
	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hello, %s", in.FirstName),
	}, nil
}
