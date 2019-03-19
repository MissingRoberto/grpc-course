package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jszroberto/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("could not connect %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorClient(conn)
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorClient) {
	req := calculatorpb.SumRequest{
		Operation: &calculatorpb.Operation{
			First:  5,
			Second: 8,
		},
	}
	resp, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Printf("error while calling sum: %v", err)
	}

	log.Printf("Result of sum: %d", resp.Result)
}
