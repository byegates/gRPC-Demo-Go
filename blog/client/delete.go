package main

import (
	"context"
	"log"

	pb "github.com/byegates/gRPC-go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	tag0 := tag + "[D] "
	log.Printf("%v[Invoked]\n", tag0)

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Printf("%v[Error] %v\n\n", tag0, err)
		return
	}

	log.Printf("%v[SUCESS] %v\n\n", tag0, id)
}
