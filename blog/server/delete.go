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
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	tag0 := tag + "[D] "
	log.Printf("%v[Invoked]\n\n", tag0)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("%v[parse ID] %v", tag0, err))
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v[Failed] with ID: %v", tag0, oid))
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("%v[Not Found] with ID: %v", tag0, oid))
	}

	return &emptypb.Empty{}, nil
}
