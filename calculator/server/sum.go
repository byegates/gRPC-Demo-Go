package main

import (
	"context"
	"log"

	pb "github.com/byegates/gRPC-go/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("%s[Invoked] {%v}\n\n", tag+"[Sum] ", in)
	return &pb.SumResponse{
		Result: in.A + in.B,
	}, nil
}
