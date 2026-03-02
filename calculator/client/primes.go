package main

import(
	"context"
	"io"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/calculator/proto"
)

func doPrimes(client pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := client.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Primes RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Response from Primes: %v", res.Prime)
	}
}
