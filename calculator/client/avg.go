package main

import (
	"context"
	"log"

	pb "github.com/byegates/gRPC-go/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	tag0 := tag + "[Bi] "
	log.Printf("%v[Invoked]\n", tag0)

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("%vError: %v\n", tag0, err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, x := range numbers {
		log.Printf("%vSending in stream: %d\n", tag0, x)
		stream.Send(&pb.AvgRequest{X: x})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("%vError while receiving response: %v\n", tag0, err)
	}

	log.Printf("%v[Recv]: %v\n\n", tag0, res.Result)
}
