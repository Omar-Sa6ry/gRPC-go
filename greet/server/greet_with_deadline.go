package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Omar-Sa6ry/grpc-go/greet/proto"
)

func (s *Server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Println("GreetWithDeadline was invoked")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Deadline exceeded")
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}
