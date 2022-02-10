package main

import (
	"context"
	"fmt"
	"gRPC/greet/Unary/GreetWithDeadline"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type exampleserver struct{}

func (*exampleserver) GreetWithDealine(ctx context.Context, in *GreetWithDeadline.GreetWithDealineRequest) (*GreetWithDeadline.GreetWithDeadlineResponse, error) {
	fmt.Println("Greet function was invoked", in)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("The client canceled the request")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		time.Sleep(time.Second * 1)
	}

	firstName := in.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &GreetWithDeadline.GreetWithDeadlineResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("hello this is the server for the RPC with deadlines")

	lis, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	GreetWithDeadline.RegisterGreetServiceServer(s, &exampleserver{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server : %v", err)
	}
}
