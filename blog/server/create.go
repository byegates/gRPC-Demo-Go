package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/byegates/gRPC-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	tag0 := tag + "[C] "
	log.Printf("%v[Invoked]\n\n", tag0)

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("%v[Interval Error] %v\n\n", tag0, err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%vFail to convert to oid\n\n", tag0))
	}

	return &pb.BlogId{Id: oid.Hex()}, nil
}
