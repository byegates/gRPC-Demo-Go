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

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	tag0 := tag + "[U] "
	log.Printf("%v[Invoked]\n\n", tag0)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("%v[parse ID] %v\n\n", tag0, err))
	}

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("%v[Fail to Update] for ID %v\n\n", tag0, err))
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("%v [Blog Not Found] with ID: %v\n\n", tag0, oid))
	}

	return &emptypb.Empty{}, nil
}
