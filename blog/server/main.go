package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Omar-Sa6ry/grpc-go/blog/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "localhost:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcServer, &Server{})

	log.Printf("Listening on %s\n", addr)
	if err = grpcServer.Serve(conn); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}