package main

import (
	"context"
	"log"

	pb "github.com/MochamadAkbar/go-grpc-example/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 3,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %d", res.Result)
}
