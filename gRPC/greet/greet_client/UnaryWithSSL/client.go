package main

import (
	"context"
	"fmt"
	"gRPC/greet/Unary/GreetAPI"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("Hello i a client")

	tls := true

	opts := grpc.WithInsecure()

	if tls {
		certFile := "D:/Udemy/go/gRPC/ssl/ca.crt" // CA trust certificate

		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificates : %v", sslErr)
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	ls, err := grpc.Dial("localhost:5000", opts)
	if err != nil {
		log.Fatalf("Failed to dial to localhost:5000 with error %v", err)
	}

	// makes sure the client connection is closed at the end
	defer ls.Close()

	c := GreetAPI.NewGreetServiceClient(ls)

	// fmt.Printf("Created client %f", c)
	doUnary(c)

}

func doUnary(c GreetAPI.GreetServiceClient) {

	fmt.Println("Starting to perform a Unary RPC")

	req := &GreetAPI.GreetRequest{
		Greeting: &GreetAPI.Greeting{
			FirstName:  "Hitesh Kumar",
			SecondName: "G Balegar",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling greet RPC :%v", err)
	}
	log.Printf("Response from Greet : %v", res)

}
