package main

import (
	"context"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/blog/proto"
)

func deleteBlog(client pb.BlogServiceClient, id string) {
	log.Println("DeleteBlog was invoked with %v", client)

	req := &pb.BlogId{
		Id: id,
	}

	_, err := client.DeleteBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to delete blog: %v", err)
	}

	log.Println("Blog has been deleted")
}