package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/byegates/gRPC-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	tag0 := tag + "[L] "
	log.Printf("%v[Invoked]\n\n", tag0)

	cur, err := collection.Find(context.Background(), primitive.D{})

	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("%v[Find] %v\n\n", tag0, err))
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}

		if err := cur.Decode(data); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("%v[Error] Decoding data from MongoDB: %v\n\n", tag0, err))
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("%v[Unchecked Internal Error] %v\n\n", tag0, err))
	}

	return nil
}
