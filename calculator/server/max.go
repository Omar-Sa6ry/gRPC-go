package main

import(
	"io"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked with %v\n", stream)

	var max int32 = 0

	for {
		req, err:=stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Receiving number: %v", req.Number)
		if req.Number > max {
			max = req.Number
			err := stream.Send(&pb.MaxResponse{
				Max: max,
			})
			
			if err != nil {
				log.Fatalf("Error while sending data to client : %v", err)
			}
		}
	}
}