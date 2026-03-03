package main

import(
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "github.com/Omar-Sa6ry/grpc-go/calculator/proto"
)

func doSqrt(client pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	req := &pb.SqrtRequest{
		Number: n,
	}

	res, err := client.Sqrt(context.Background(), req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Fatalf("Error message from server: %v", e.Message())
			log.Fatalf("Error code from server: %v", e.Code())
		
			if e.Code() == codes.InvalidArgument {
				log.Fatalf("We probably sent a negative number: %v", e.Message())
				return
			}
		} else {
			log.Fatalf("non-RPC error: %v", err)
		}

		log.Fatalf("Error while calling Sqrt RPC: %v", err)
	}

	log.Printf("Response from Sqrt: %v", res.Result)
}