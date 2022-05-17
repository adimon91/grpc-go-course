package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/adimon91/grpc-go-course/greet/proto"
)

const (
	addr string = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	// close the connection
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
}
