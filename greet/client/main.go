package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/MochamadAkbar/go-grpc-example/greet/proto"
)

var Arr = "localhost:50051"

func main() {
	// for development use => grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(Arr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	// doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	doGreetWithDeadline(c, 1*time.Second)
}
