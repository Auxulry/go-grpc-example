package main

import (
	"log"
	"net"

	pb "github.com/MochamadAkbar/go-grpc-example/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.CalculatorServiceServer
}

var Addr = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", Addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %s\n", Addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
