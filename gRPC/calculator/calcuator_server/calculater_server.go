package main

import (
	"context"
	"fmt"
	"gRPC/calculator/calculator_proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type calulateServer struct{}

func (*calulateServer) CalcuteSum(ctx context.Context, in *calculator_proto.CalculatorServerRequest) (*calculator_proto.CalculatorServerResponse, error) {
	//(*CalculatorServerResponse, error)
	fmt.Println("Calculate function was invoked on the server ", ctx)

	first_number := in.GetCalculate().GetFirstNumber()
	second_number := in.GetCalculate().GetSecondNumer()

	sum := first_number + second_number

	fmt.Printf("The sum of the first %v and the second number %v is %v", first_number, second_number, sum)

	res := &calculator_proto.CalculatorServerResponse{
		ResultSum: sum,
	}

	return res, nil

}
func main() {

	fmt.Println("Hello welcome to server service")

	lis, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()

	calculator_proto.RegisterCalculateSumServiceServer(s, &calulateServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

}
