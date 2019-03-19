package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jszroberto/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked %v\n", req)
	result := req.GetOperation().GetFirst() + req.GetOperation().GetSecond()
	return &calculatorpb.SumResponse{
		Result: result,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Printf("cannot listen %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
