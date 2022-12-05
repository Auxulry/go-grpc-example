package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/MochamadAkbar/go-grpc-example/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while calling doMax: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		nums := []int32{1, 5, 3, 6, 2, 12}
		for _, req := range nums {
			log.Printf("Sending req: %v\n", req)

			stream.Send(&pb.MaxRequest{
				Number: req,
			})

			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
