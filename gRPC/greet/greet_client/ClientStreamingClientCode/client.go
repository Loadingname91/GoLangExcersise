package main

import (
	"context"
	"fmt"
	"gRPC/greet/Stream/GreetAPIClientManyTimes"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello Im a client")

	ls, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial to localhost:5000 with error %v", err)
	}

	// makes sure the client connection is closed at the end
	defer ls.Close()

	c := GreetAPIClientManyTimes.NewGreetServiceClient(ls)

	// fmt.Printf("Created client %f", c)
	doLongClientStreaming(c)

}

func doLongClientStreaming(c GreetAPIClientManyTimes.GreetServiceClient) {

	fmt.Println("Starting up client streaming")

	requests := []*GreetAPIClientManyTimes.LongGreetRequest{
		{
			Greeting: &GreetAPIClientManyTimes.Greeting{
				FirstName: "HiteshKumar",
			},
		},
		{
			Greeting: &GreetAPIClientManyTimes.Greeting{
				FirstName: "ShilajitRoy",
			},
		},
		{
			Greeting: &GreetAPIClientManyTimes.Greeting{
				FirstName: "pakshalBhandari",
			},
		},
		{
			Greeting: &GreetAPIClientManyTimes.Greeting{
				FirstName: "Balegar",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling the stream GreetManyTimes %v\n", err)
	}
	for _, req := range requests {
		fmt.Printf("Sending req %v\n", req)
		stream.Send(req)
		time.Sleep(time.Nanosecond * 1000)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while streaming the data from long greet %v\n", err)
	}

	fmt.Printf("Received data from the server %v\n", res)

}
