package main

import (
	"log"

	pb "github.com/MochamadAkbar/go-grpc-example/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id) // valid
	// readBlog(c, "aNonExistingId") // invalid
	updateBlog(c, id)
	ListBlogs(c)
	deleteBlog(c, id)
}
