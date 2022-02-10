package main

import (
	"context"
	"fmt"
	"gRPC/greet/Unary/GreetAPI"
	"log"
	"net"

	"google.golang.org/grpc"
)

type exampleserver struct{}

func (*exampleserver) Greet(ctx context.Context, in *GreetAPI.GreetRequest) (*GreetAPI.GreetResponse, error) {
	fmt.Println("Greet function was invoked", in)
	firstName := in.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &GreetAPI.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	GreetAPI.RegisterGreetServiceServer(s, &exampleserver{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server : %v", err)
	}
}
