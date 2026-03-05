package main

import (
	pb "github.com/Omar-Sa6ry/grpc-go/blog/proto"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Server) DeleteBlog(ctx context.Context, req *pb.BlogId) (*emptypb.Empty, error) {
	log.Println("DeleteBlog was invoked with %v", req)

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid ID")
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete blog")
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Cannot find blog")
	}

	return &emptypb.Empty{}, nil
}