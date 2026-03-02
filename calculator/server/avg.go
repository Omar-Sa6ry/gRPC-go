package main

import(
	"io"
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked with %v\n", stream)

	var sum int32 = 0
	var count int32 = 0

	for {
		req,err:=stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Receiving number: %v", req.Number)
		sum+=req.Number
		count++
	}
}