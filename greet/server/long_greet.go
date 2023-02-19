package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/byegates/gRPC-go/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	tag0 := tag + "[CS] "
	log.Printf("%v[Invoked]\n", tag0)

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Printf("\n\n")
			return stream.SendAndClose(
				&pb.GreetResponse{
					Result: res,
				},
			)
		}

		if err != nil {
			log.Fatalf("%vError while reading client stream: %v\n", tag0, err)
		}

		log.Printf("%v[Recv]: %v\n", tag0, req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
