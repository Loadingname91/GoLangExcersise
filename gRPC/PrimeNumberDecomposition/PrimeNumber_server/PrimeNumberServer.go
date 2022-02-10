package main

import (
	"fmt"
	"gRPC/PrimeNumberDecomposition/PrimeNumberPb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeNumberDecomposition(req *PrimeNumberPb.PrimeNumberRequest, stream PrimeNumberPb.PrimeNumberService_PrimeNumberDecompositionServer) error {

	fmt.Println("Starting up the stream to send the prime number decomposition with req ", req)
	GivenNumber := req.GetNumberGiven().GetGivenNumber()

	k := int32(2)
	N := GivenNumber
	for N > 1 {
		if int32((N)%k) == 0 {
			result := k
			res := &PrimeNumberPb.PrimeNumberResponse{
				Result: result,
			}
			stream.Send(res)
			N = N / k
		} else {
			k = k + 1
		}
	}
	return nil

}

func main() {
	fmt.Println("Starting up the server for now")

	lis, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("Couldnt not listen to localhost 5000 %v", err)
	}

	s := grpc.NewServer()

	PrimeNumberPb.RegisterPrimeNumberServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to connect to server %v", err)
	}

}
