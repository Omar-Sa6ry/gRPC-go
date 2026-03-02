package main

import(
	"log"

	pb "github.com/Omar-Sa6ry/grpc-go/calculator/proto"
)

func (s *Server) Primes(req *pb.PrimeRequest,stream pb.CalculatorService_PrimesServer) error {
	log.Println("Primes function was invoked with %v\n", req)

	number:= req.Number
	divisor:=int64(2)

	for number > 1 {
		if number % divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Prime: divisor,
			})
			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}