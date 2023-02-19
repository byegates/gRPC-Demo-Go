package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/byegates/gRPC-go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	tag0 := tag + "[Bi] "
	log.Printf("%v[Invoked]\n", tag0)

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("%vError while creating stream: %v\n", tag0, err)
	}

	waitc := make(chan struct{})

	// stream [Send]
	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, x := range numbers {
			log.Printf("%v[Send] {%v}\n", tag0, x)
			stream.Send(&pb.MaxRequest{X: x})
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
		log.Printf("%v[Send] [End]\n\n", tag0)
	}()

	// stream [Recv]
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("%v[Recv] [Error] from stream: %v\n", tag0, err)
				break
			}

			log.Printf("%v[Recv] [max] %v\n", tag0, res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
