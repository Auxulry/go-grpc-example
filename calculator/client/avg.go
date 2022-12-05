package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MochamadAkbar/go-grpc-example/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	reqs := []*pb.AverageRequest{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Result)
}
