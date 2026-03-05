package main

import (
	"context"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/blog/proto"
)

func updateBlog(client pb.BlogServiceClient, id string) {
	log.Println("UpdateBlog was invoked with %v", client)

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "123",
		Title:    "My Updated Blog",
		Content:  "This is the updated content of my blog.",
	}
	
	_, err := client.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Failed to update blog: %v", err)
	}

	log.Println("Blog has been updated")
}