package main

import (
	"fmt"
	"gRPC/greet/Stream/GreetAPIManyTimes"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetManyTimes(req *GreetAPIManyTimes.GreetRequestManyTimes, stream GreetAPIManyTimes.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello" + firstName + "Number" + strconv.Itoa(i)
		res := &GreetAPIManyTimes.GreetResponseManyTimes{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {

	fmt.Println("hello starting up the server for the stream")

	lis, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("Failed to listen to localhost :5000 %v", err)
	}

	s := grpc.NewServer()

	GreetAPIManyTimes.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to connect to server %v", err)
	}

}
