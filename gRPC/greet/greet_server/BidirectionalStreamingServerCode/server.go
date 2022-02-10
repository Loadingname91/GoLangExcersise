package main

import (
	"fmt"
	"gRPC/greet/Stream/BiDirectionalAPI"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetEveryone(stream BiDirectionalAPI.GreetService_GreetEveryoneServer) error {

	fmt.Printf("GreetEveryone function was invoked with a streaming request\n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "

		sendErr := stream.Send(&BiDirectionalAPI.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}

}
func main() {

	fmt.Println("hello starting up the server for Bi directional streaming")

	lis, err := net.Listen("tcp", "0.0.0.0:5000")

	if err != nil {
		log.Fatalf("Failed to connect to localhost :5000 with error %v", err)
	}

	s := grpc.NewServer()

	BiDirectionalAPI.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server with error %v", err)
	}

}
