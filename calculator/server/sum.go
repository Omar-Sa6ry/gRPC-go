package main

import (
	"context"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("Sum function was invoked with %v\n", req)
	return &pb.SumResponse{
		Result: req.FirstNumber + req.SecondNumber,
	}, nil
}