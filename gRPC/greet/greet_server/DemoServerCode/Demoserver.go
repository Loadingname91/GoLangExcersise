package main

import (
	"fmt"
	"gRPC/greet/PbDemo"
	"log"
	"net"

	"google.golang.org/grpc"
)

type exampleserver struct{}

func main() {
	fmt.Println("hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:5005")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	exampleserver := PbDemo.UnimplementedGreetServiceServer{}
	PbDemo.RegisterGreetServiceServer(s, &exampleserver)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server : %v", err)
	}

}
