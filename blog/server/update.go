package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Omar-Sa6ry/grpc-go/blog/proto"
)

func (s *Server) UpdateBlog(ctx context.Context, req *pb.Blog) (*emptypb.Empty, error) {
	log.Println("UpdateBlog was invoked with %v", req)

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid ID")
	}

	data := &BlogItem{
		AuthorID: req.AuthorId,
		Title:    req.Title,
		Content:  req.Content,
	}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": data})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update blog")
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Cannot find blog")
	}

	return &emptypb.Empty{}, nil
	
}