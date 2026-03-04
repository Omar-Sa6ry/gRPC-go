package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/blog/proto"
)

func createBlog(client pb.BlogServiceClient) string {
	log.Println("CreateBlog was invoked with %v", client)
	blog := &pb.Blog{
		AuthorId: "123",
		Title:    "My First Blog",
		Content:  "This is the content of my first blog.",
	}

	resp, err := client.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Failed to create blog: %v", err)
	}

	fmt.Println("Blog has been created:", resp.Id)
	return resp.Id
}