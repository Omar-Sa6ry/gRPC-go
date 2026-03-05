package main

import(
	"context"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Omar-Sa6ry/grpc-go/blog/proto"
)

func listBlogs(client pb.BlogServiceClient) {
	log.Println("ListBlogs was invoked with %v", client)

	stream, err := client.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Failed to list blogs: %v", err)
	}

	for {
		blog, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to receive blog: %v", err)
		}
		
		log.Printf("Blog received: %v", blog)
	}
}