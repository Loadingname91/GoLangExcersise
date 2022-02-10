package main

import (
	"context"
	"fmt"
	"gRPC/calculateSqaureRoot/calculate_proto"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) SqaureRoot(ctx context.Context, in *calculate_proto.SquareRootRequest) (*calculate_proto.SqaureRootResponse, error) {

	fmt.Println("The function SqaureRoot is starting up")
	number := in.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Recieved a negative number %v", number),
		)
	}
	return &calculate_proto.SqaureRootResponse{
		Result: math.Sqrt(float64(number)),
	}, nil

}
func main() {

	fmt.Println("This the server code starting up")

	lis, err := net.Listen("tcp", "0.0.0.0:5000")

	if err != nil {
		log.Fatalf("Error while starting up the server on port 5000 %v", err)
	}

	s := grpc.NewServer()

	calculate_proto.RegisterFindSquareRootServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

}
