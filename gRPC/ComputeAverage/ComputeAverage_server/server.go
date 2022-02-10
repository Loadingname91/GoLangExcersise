package main

import (
	"fmt"
	"gRPC/ComputeAverage/ComputeAveragePb"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) ComputeAverage(stream ComputeAveragePb.ComputedAverage_ComputeAverageServer) error {
	fmt.Println("LongGreet function was invoked with a steaming request")
	numbers := []float32{}
	for {
		req, err := stream.Recv()

		if err == io.EOF {

			sum := float32(0)
			for _, val := range numbers {
				sum += val
			}

			result := sum / float32(len(numbers))

			errStream := stream.SendAndClose(&ComputeAveragePb.ComputedAverageResponse{
				Result: result,
			})
			return errStream
		}

		if err != nil {
			log.Fatalf("Failed in reading the stream with the error %v", err)
		}

		input := req.GetInput().GetNumber()
		numbers = append(numbers, input)
	}

}
func main() {

	fmt.Println("hello starting up the server for client streaming")

	lis, err := net.Listen("tcp", "0.0.0.0:5000")

	if err != nil {
		log.Fatalf("Failed to connect to localhost :5000 with error %v", err)
	}

	s := grpc.NewServer()

	ComputeAveragePb.RegisterComputedAverageServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server with error %v", err)
	}

}
