package main

import (
	"context"
	"log"

	pb "github.com/byegates/gRPC-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	tag0 := tag + "[C] "
	log.Printf("%v[Invoked]\n", tag0)

	blog := &pb.Blog{
		AuthorId: "肖骁",
		Title:    "I can I bb",
		Content:  "你们希望我做一个good boy, 我只想让自己当一个bad girl",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("%v[Error] %v\n\n", tag0, err)
	}

	log.Printf("%v[Created] %v\n\n", tag0, res.Id)
	return res.Id
}
