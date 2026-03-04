package main

import(
	"context"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/blog/proto"
)

func readBlog(client pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("ReadBlog was invoked with %v", client)
	
	req := &pb.BlogId{
		Id: id,
	}

	res, err := client.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to read blog: %v", err)
	}

	log.Printf("Blog read: %v", res)

	return res
}