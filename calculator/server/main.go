package main

import (
	"log"
	"net"

	pb "github.com/Omar-Sa6ry/grpc-go/calculator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "localhost:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(grpcServer, &Server{})

	reflection.Register(grpcServer)

	log.Printf("Listening on %s\n", addr)
	if err = grpcServer.Serve(conn); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}