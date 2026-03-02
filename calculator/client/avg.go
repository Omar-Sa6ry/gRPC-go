package main

import(
	"context"
	"log"
	"time"

	pb "github.com/Omar-Sa6ry/grpc-go/calculator/proto"
)

func doAvg(client pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")


	stream, err := client.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg RPC: %v", err)
	}

	numbers := []int32{3, 5, 54, 23}

	for _,num:=range numbers {
		stream.Send(&pb.AvgRequest{
			Number: num,
		})

		log.Printf("Sending: %v", num)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving: %v", err)
	}

	log.Printf("Avg: %v", res.Result)
}