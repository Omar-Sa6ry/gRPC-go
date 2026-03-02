package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Omar-Sa6ry/grpc-go/greet/proto"
)

func doGreetEveryone(client pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")


	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetEveryone RPC: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Omar"},
		{FirstName: "Ahmed"},
		{FirstName: "Sabry"},
	}

	watic := make(chan struct{})

	go func() {
		for _,req:= range reqs {
			stream.Send(req)
			log.Printf("Sending: %v", req)
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

		log.Printf("Received: %v", res.Result)
	}

	close(watic)
	}()

	<-watic

}