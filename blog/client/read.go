package main

import (
	"context"
	"log"

	pb "github.com/byegates/gRPC-go/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	tag0 := tag + "[R] "
	log.Printf("%v[Invoked]\n", tag0)

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("%v[Error] %v\n\n", tag0, err)
		return nil
	}

	log.Printf("%v[Blog] {%v}\n\n", tag0, res)
	return res
}
