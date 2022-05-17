package main

import (
	"context"
	"log"

	pb "github.com/adimon91/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.CalculatorRequest{
		FirstNumber:  5,
		SecondNumber: 5,
	})

	if err != nil {
		log.Fatalf("Failed to sum: %v\n", err)
	}

	log.Printf("Sum result is: %d\n", res.Result)
}
