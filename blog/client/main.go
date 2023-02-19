package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/byegates/gRPC-go/blog/proto"
)

var addr string = "localhost:50054"
var tag string = "[Client] "

func main() {
	con, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("%vFailed to connect to '%v' with Error: %v\n", tag, addr, err)
	}

	defer con.Close()

	c := pb.NewBlogServiceClient(con)

	id := createBlog(c)
	readBlog(c, "0000")                     // negative case
	readBlog(c, id)                         // postive case
	readBlog(c, "63f1714ce2d734e4c8a193dd") // negative case
	updateBlog(c, id)
	readBlog(c, id)
	deleteBlog(c, id)
	ListBlogs(c)
	//...
}
