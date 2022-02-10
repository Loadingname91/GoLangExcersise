package main

import (
	"context"
	"fmt"
	"gRPC/greet/Unary/GreetWithDeadline"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hi this is a client")

	ls, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to dial to localhost : 5000 with error %v", err)
	}

	defer ls.Close()

	c := GreetWithDeadline.NewGreetServiceClient(ls)

	doUnaryCallWithDeadline(c, 1*time.Second)

	doUnaryCallWithDeadline(c, 10*time.Second)

}

func doUnaryCallWithDeadline(c GreetWithDeadline.GreetServiceClient, timeout time.Duration) {

	fmt.Println("Starting to perform a Unary RPC")

	req := &GreetWithDeadline.GreetWithDealineRequest{
		Greeting: &GreetWithDeadline.Greeting{
			FirstName:  "Hitesh Kumar",
			SecondName: "G Balegar",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	res, err := c.GreetWithDealine(ctx, req)

	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Time out was hit ! Deadline was exceeded")
			} else {
				fmt.Println("Unexpected error ", statusErr)
			}
		} else {
			fmt.Println("Error while calling the GreetWithDeadline RPC", err)
		}
		return
	}

	log.Print("Response from the GreetWithDeadline", res.Result)

}
