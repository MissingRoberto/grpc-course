package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jszroberto/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Hello I'm a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Making a Unary request")
	req := greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Roberto",
			LastName:  "Jimenez",
		},
	}

	resp, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Printf("error while calling greet RPC: &v", err)
	}

	log.Printf("Response from greet %v\n", resp.GetResult())
}
