package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MochamadAkbar/go-grpc-example/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoke")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Akbar",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exeeded !")
				return
			} else {
				log.Fatalf("unexpected gRPC error: %v", err)
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
