package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/byegates/gRPC-go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	tag0 := tag + "[Sqrt] "
	log.Printf("%s[Invoked]\n\n", tag0)

	if in.X < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Got negative number: %d", in.X))
	}

	return &pb.SqrtResponse{Val: math.Sqrt(float64(in.X))}, nil
}
