package main

import (
	"context"
	"fmt"
	"gRPC/greet/Unary/GreetAPI"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	tls := true

	opts := []grpc.ServerOption{}

	if tls {
		certFile := "D:/Udemy/go/gRPC/ssl/server.crt"
		keyFile := "D:/Udemy/go/gRPC/ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)

		if sslErr != nil {
			log.Fatalf("Failed loading certificates : %v", sslErr)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	GreetAPI.RegisterGreetServiceServer(s, &exampleserver{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server : %v", err)
	}
}
