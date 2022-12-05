package main

import (
	"log"

	pb "github.com/MochamadAkbar/go-grpc-example/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with: %v\n", in)

	number := in.Number
	divisior := int64(2)

	for number > 1 {
		if number%divisior == 0 {
			stream.Send(&pb.PrimesResponse{
				Result: divisior,
			})

			number /= divisior
		} else {
			divisior++
		}
	}

	return nil
}
