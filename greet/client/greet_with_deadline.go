package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Omar-Sa6ry/grpc-go/greet/proto"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Omar",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
				return
			} else{
				log.Printf("Unexpected error: %v\n", e.Code())
			}
		} else{
			log.Fatalf("a non grpc error: %v\n", err)
		}
		return
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}