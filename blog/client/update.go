package main

import (
	"context"
	"log"

	pb "github.com/byegates/gRPC-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	tag0 := tag + "[U] "
	log.Printf("%v[Invoked]\n", tag0)

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "肖骁2",
		Title:    "I can I bb2",
		Content:  "你们希望我做一个good boy, 我只想让自己当一个bad girl -- Updated",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Printf("%v[Error] %v\n\n", tag0, err)
		return
	}

	log.Printf("%v[SUCCESS]\n\n", tag0)
}
