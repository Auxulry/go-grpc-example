package main

import (
	"context"
	"log"

	pb "github.com/MochamadAkbar/go-grpc-example/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error Code from server: %d\n", e.Code())
			log.Printf("Error Message from server: %s\n", e.Message())

			if e.Code() == codes.InvalidArgument {
				log.Printf("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %v\n", res)
}
