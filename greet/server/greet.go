package main

import (
	"context"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/greet/proto"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("Greet function was invoked with %v\n", req)
	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}