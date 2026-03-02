package main

import(
	"context"
	"log"
	"time"

	pb "github.com/Omar-Sa6ry/grpc-go/greet/proto"
)

func doLongGreet(client pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Omar"},
		{FirstName: "Ahmed"},
		{FirstName: "Sabry"},
	}

	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet RPC: %v", err)
	}

	for _,req:= range reqs {
		stream.Send(req)
		log.Printf("Sending: %v", req)
		time.Sleep(1 * time.Second)
	}

	res,err:=stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving: %v", err)
	}

	log.Printf("Response from LongGreet: %v", res.Result)
}