package main

import (
	"context"
	"fmt"
	"gRPC/calculateSqaureRoot/calculate_proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello this is the client for calculating the sqaure root of a number")

	ls, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error while dialing up the connection to localhost:5000 %v", err)
	}

	defer ls.Close()

	c := calculate_proto.NewFindSquareRootClient(ls)

	PerformUnaryOperation(c)

}

func PerformUnaryOperation(c calculate_proto.FindSquareRootClient) {

	fmt.Println("Starting the function to find the Sqaure Root of a given number using Unary")

	number := int32(10)
	// Correct Request
	doErrorCall(c, number)

	// Invalid Request
	number = -12
	doErrorCall(c, number)

}

func doErrorCall(c calculate_proto.FindSquareRootClient, number int32) {
	res, err := c.SqaureRoot(context.Background(), &calculate_proto.SquareRootRequest{
		Number: number,
	})

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			log.Println(respErr.Message())
			log.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably just sent a negative number")
				return
			}
		} else {
			log.Fatal("Received a big error from by calling the sqaureRoot function", err)
			return
		}
	}
	fmt.Printf("Result of sqaure root of %v is %v", number, res.GetResult())

}
