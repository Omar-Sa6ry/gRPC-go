package main

import(
	"context"
	"log"
	"math"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Omar-Sa6ry/grpc-go/calculator/proto"
)

func (s *Server) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Println("Sqrt function was invoked with %v\n", req)

	if req.Number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Number must be non-negative: %d", req.Number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(req.Number)),
	}, nil
}