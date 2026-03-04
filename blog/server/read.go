package main

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	pb "github.com/Omar-Sa6ry/grpc-go/blog/proto"
)

func (s *Server) ReadBlog(ctx context.Context, req *pb.BlogId) (*pb.Blog, error) {
	log.Println("ReadBlog was invoked with %v", req)

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid ID")
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	result := collection.FindOne(ctx, filter)
	if err := result.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, "Cannot find blog")
	}

	return documentToBlog(data), nil
}