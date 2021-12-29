package main

import (
	"fmt"
	"gRPC/greet/PbDemo"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i a client")

	ls, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial to localhost:5000 with error %v", err)
	}

	// makes sure the client connection is closed at the end
	defer ls.Close()

	c := PbDemo.NewGreetServiceClient(ls)

	fmt.Printf("Created client %f", c)

}
