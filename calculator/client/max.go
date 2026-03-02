package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Omar-Sa6ry/grpc-go/calculator/proto"
)

func doMax(client pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Max RPC: %v", err)
	}

	watic := make(chan struct{})

	go func() {
		numbers := []int32{3, 5, 54, 23}
		for _, num := range numbers {
			stream.Send(&pb.MaxRequest{
				Number: num,
			})

			log.Printf("Sending: %v", num)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
		close(watic)
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
			}

			log.Printf("Received a new maximum: %v", res.Max)
		}

		close(watic)
	}()

	<-watic
}