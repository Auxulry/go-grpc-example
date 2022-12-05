package main

import (
	"io"
	"log"

	pb "github.com/MochamadAkbar/go-grpc-example/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("AVG was invoked")

	res := int64(0)
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(res) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving req: %v\n", req)
		res += req.Number
		count++
	}
}
