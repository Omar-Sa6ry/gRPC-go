package main

import(
	"context"
	"log"
	
	pb "github.com/Omar-Sa6ry/grpc-go/greet/proto"
)

func doGreet(client pb.GreetServiceClient){
	log.Println("doGreet was invoked")
	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Omar",
	})

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}