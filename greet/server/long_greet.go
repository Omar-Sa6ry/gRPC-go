package main

import(
	"fmt"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked with %v\n", stream)

	res:=""
	for {
		req,err:=stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		res+=fmt.Sprintf("Hello %s\n",req.FirstName)
	}

}