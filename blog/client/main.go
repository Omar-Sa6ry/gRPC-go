package main

import (
	"log"

	"google.golang.org/grpc"

	pb "github.com/Omar-Sa6ry/grpc-go/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)

	id :=createBlog(client)
	readBlog(client, id)
	// readBlog(client, "ifg")
	updateBlog(client, id)
	listBlogs(client)
	deleteBlog(client, id)
}
