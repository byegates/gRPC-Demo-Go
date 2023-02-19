package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/byegates/gRPC-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	tag0 := tag + "[R] "
	log.Printf("%v[Invoked]\n\n", tag0)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("[Fail to parse ID: '%v'] [%v]", in.Id, err))
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("[ID: %v] [%v]", oid, err))
	}

	return documentToBlog(data), nil
}
