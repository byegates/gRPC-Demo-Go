package main

import (
	"context"
	"log"

	pb "github.com/byegates/gRPC-go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	tag0 := tag + "[Sqrt] "
	log.Printf("%v[Invoked]\n", tag0)

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{X: n})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("%v[Error] [msg]:%v\n", tag0, e.Message())
			log.Printf("%v[Error] [code]:%v\n", tag0, e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("%vWe probably sent a negative number: %v\n\n", tag0, n)
				return
			}
		} else {
			log.Fatalf("%vA non gPRC error: %v\n", tag0, err)
		}
	}

	log.Printf("%v[Recv]: Sqrt(%v) = %v\n\n", tag0, n, res.Val)
}
