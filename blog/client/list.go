package main

import (
	"context"
	"io"
	"log"

	pb "github.com/byegates/gRPC-go/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ListBlogs(c pb.BlogServiceClient) {
	tag0 := tag + "[L] "
	log.Printf("%v[Invoked]\n", tag0)

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Printf("%v[Error] Creating Stream: %v\n\n", tag0, err)
		return
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			log.Printf("\n")
			break
		}

		if err != nil {
			log.Printf("%v[Error] [Recv] from Stream: %v\n\n", tag0, err)
			return
		}

		log.Printf("%v[Blog] {%v}\n", tag0, res)
	}

}
