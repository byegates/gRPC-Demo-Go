package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/byegates/gRPC-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	tag0 := tag + "[DL] "
	log.Printf("%v[Invoked] : {%v}\n\n", tag0, in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			msg := fmt.Sprintf("%v[Err] Client canceled request\n\n", tag0)
			log.Print(msg)
			return nil, status.Error(codes.Canceled, msg)
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hello, %s", in.FirstName),
	}, nil

}
