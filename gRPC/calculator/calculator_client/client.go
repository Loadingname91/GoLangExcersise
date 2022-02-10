package main

import (
	"context"
	"fmt"
	"gRPC/calculator/calculator_proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello this is the calculator client")

	ls, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to localhost 5000 with error : %v", err)
	}

	defer ls.Close()

	c := calculator_proto.NewCalculateSumServiceClient(ls)

	// Perform Sum

	PerformSum(c)

}

func PerformSum(c calculator_proto.CalculateSumServiceClient) {

	fmt.Println("Starting to perform a sum using Unary Apis")

	req := &calculator_proto.CalculatorServerRequest{
		Calculate: &calculator_proto.Calculator{
			FirstNumber: 50,
			SecondNumer: 40,
		},
	}

	res, err := c.CalcuteSum(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling calcuteSum : %v", res)
	}

	log.Printf("Response from calcuteSum with result : %v", res)

}
