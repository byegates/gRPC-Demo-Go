package main

import (
	"context"
	"log"
	"time"

	pb "github.com/byegates/gRPC-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, dl time.Duration) {
	tag0 := tag + "[DL] "
	log.Printf("%v[Invoked]\n", tag0)

	ctx, cancel := context.WithTimeout(context.Background(), dl)
	defer cancel()

	req := &pb.GreetRequest{FirstName: "邱晨"}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("%vDeadline exceeded!\n\n", tag0)
				return
			} else {
				log.Fatalf("%vAn unexpected gPRC error: %v\n\n", tag0, err)
			}
		} else {
			log.Fatalf("%vA non gPRC error: %v\n\n", tag0, err)
		}
	}

	log.Printf("%v[Recv] %v\n\n", tag0, res.Result)
}
