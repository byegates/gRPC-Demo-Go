package main

import (
	"io"
	"log"

	pb "github.com/byegates/gRPC-go/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	tag0 := tag + "[BI] "
	log.Printf("%v[Invoked]\n", tag0)

	var max int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Printf("%v[End]\n\n", tag0)
			return nil
		}

		if err != nil {
			log.Fatalf("%v[Error] receiving from stream: %v\n", tag0, err)
		}

		if max < req.X {
			max = req.X
		}

		err = stream.Send(&pb.MaxResponse{Result: max})

		if err != nil {
			log.Fatalf("%v[Error] [Send] to stream: %v\n", tag0, err)
		}
	}
}
