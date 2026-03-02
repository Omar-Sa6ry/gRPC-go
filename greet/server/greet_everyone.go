package main

import(
	"fmt"
	"io"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked with %v\n", stream)

	for {
		req,err:=stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		res := fmt.Sprintf("Hello %s\n",req.FirstName)
		log.Printf("Receiving: %v", req)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client : %v", err)
		}
	}
}
