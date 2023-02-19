package main

import (
	"context"
	"log"

	pb "github.com/byegates/gRPC-go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	tag0 := tag + "[sum] "
	log.Printf("%v[Invoked]\n", tag0)
	var a, b int32
	a, b = 3, 10
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: a,
		B: b,
	})

	if err != nil {
		log.Fatalf("%v[Error] %v\n", tag0, err)
	}

	log.Printf("%v[Recv]: %v + %v = %v\n\n", tag0, a, b, res.Result)
}
