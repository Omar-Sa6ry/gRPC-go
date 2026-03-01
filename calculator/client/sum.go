package main

import (
	"context"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/calculator/proto"
)

func doSum(client pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := client.Sum(context.Background(), &pb.SumRequest{
		FirstNumber: 10,
		SecondNumber: 7,
	})

	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.Result)
}