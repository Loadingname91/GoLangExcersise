package main

import (
	"fmt"
	"gRPC/FindMaxium/FindMaximumPb"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) FindMaximum(stream FindMaximumPb.ComputedAverage_FindMaximumServer) error {
	fmt.Println("Received FindMaximum RPC")
	maximum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		number := req.GetNumber()
		if number > maximum {
			maximum = number
			sendErr := stream.Send(&FindMaximumPb.FindMaximumResponse{
				Maximum: maximum,
			})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return sendErr
			}
		}
	}
}
func main() {

	fmt.Println("hello starting up the server for client streaming")

	lis, err := net.Listen("tcp", "0.0.0.0:5000")

	if err != nil {
		log.Fatalf("Failed to connect to localhost :5000 with error %v", err)
	}

	s := grpc.NewServer()

	FindMaximumPb.RegisterComputedAverageServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server with error %v", err)
	}

}
