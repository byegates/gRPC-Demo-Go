package main

import (
	"context"
	"log"
	"net"

	pb "github.com/byegates/gRPC-go/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50054"
var collection *mongo.Collection
var tag string = "[Server] "

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatalf("%vFailed to create Mongo client : %v\n", tag, err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf("%vFailed to connect to Mongo : %v\n", tag, err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("%vFailed to listen on : %v\n", tag, err)
	}

	log.Printf("%vListening on %s\n", tag, addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	// reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("%vFailed to serve: %v\n", tag, err)
	}
}
