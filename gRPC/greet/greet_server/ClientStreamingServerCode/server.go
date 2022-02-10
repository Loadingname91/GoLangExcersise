package main

import (
	"fmt"
	"gRPC/greet/Stream/GreetAPIClientManyTimes"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) LongGreet(stream GreetAPIClientManyTimes.GreetService_LongGreetServer) error {
	fmt.Println("LongGreet function was invoked with a steaming request")
	result := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			errStream := stream.SendAndClose(&GreetAPIClientManyTimes.LongGreetResponse{
				Result: result,
			})
			return errStream
		}

		if err != nil {
			log.Fatalf("Failed in reading the stream with the error %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}

}
func main() {

	fmt.Println("hello starting up the server for client streaming")

	lis, err := net.Listen("tcp", "0.0.0.0:5000")

	if err != nil {
		log.Fatalf("Failed to connect to localhost :5000 with error %v", err)
	}

	s := grpc.NewServer()

	GreetAPIClientManyTimes.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server with error %v", err)
	}

}
