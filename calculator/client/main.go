package main

import (
	"log"

	pb "github.com/adimon91/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr string = "localhost:5051"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	//close the connection
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	doSum(c)
}
