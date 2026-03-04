package main

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	pb "github.com/Omar-Sa6ry/grpc-go/blog/proto"
)

func (s *Server) CreateBlog(ctx context.Context, req *pb.Blog) (*pb.Blog, error) {
	log.Println("CreateBlog was invoked with %v", req)

	data := BlogItem{
		AuthorID: req.AuthorId,
		Title:    req.Title,
		Content:  req.Content,
	}

	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	oid,ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("Failed to get inserted ID")
	}
	
	return &pb.Blog{
		Id:       oid.Hex(),
		AuthorId: req.AuthorId,
		Title:    req.Title,
		Content:  req.Content,
	}, nil
}