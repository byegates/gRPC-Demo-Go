package main

import (
	"context"
	"io"
	"log"

	pb "github.com/byegates/gRPC-go/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	tag0 := tag + "[primes] "
	log.Printf("%v[Invoked]\n", tag0)
	x := int64(12390392840)
	stream, err := c.Primes(context.Background(), &pb.PrimeRequest{
		X: x,
	})

	if err != nil {
		log.Fatalf("%vCould not doPrimes: %v\n", tag0, err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("%vError while reading Primes stream: %v\n", tag0, err)
		}

		log.Printf("%v[Recv] %v : %v\n", tag0, x, res.Result)
	}

	log.Printf("\n\n")
}
