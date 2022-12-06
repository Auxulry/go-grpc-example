package main

import (
	"log"
	"net"

	pb "github.com/MochamadAkbar/go-grpc-example/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	pb.GreetServiceServer
}

var Addr = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", Addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %s\n", Addr)

	opts := []grpc.ServerOption{}

	tls := true

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"

		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed to loading certificates: %v\v", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
