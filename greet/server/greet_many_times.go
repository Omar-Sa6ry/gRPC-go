package main

import (
	"fmt"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		response := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)
		
		stream.Send(&pb.GreetResponse{
			Result: response,
		})
	}

	return nil
}
