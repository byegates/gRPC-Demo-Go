package main

import (
	"log"

	pb "github.com/byegates/gRPC-go/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("%s[Primes] [Invoked] {%v}\n", tag, in)

	x := in.X
	divisor := int64(2)

	for x > 1 {
		if x%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			x /= divisor
		} else {
			divisor++
		}
	}

	log.Printf("\n\n")
	return nil
}
