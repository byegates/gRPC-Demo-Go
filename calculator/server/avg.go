package main

import (
	"io"
	"log"

	pb "github.com/byegates/gRPC-go/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	tag0 := tag + "[Avg] "
	log.Printf("%s[Invoked]\n", tag0)

	var sum int32 = 0
	cnt := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Printf("%v[Exit] Respond for %d numbers\n\n", tag0, cnt)
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(cnt),
			})
		}

		if err != nil {
			log.Fatalf("%vError reading client stream: %v\n", tag0, err)
		}

		log.Printf("%v[Recv] in stream: %d\n", tag0, req.X)
		sum += req.X
		cnt++
	}
}
