package main

import (
	"context"
	"fmt"
	"gRPC/greet/Stream/BiDirectionalAPI"
	"io"
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

	c := BiDirectionalAPI.NewGreetServiceClient(ls)

	// fmt.Printf("Created client %f", c)
	doBiDirectionalStreaming(c)

}

func doBiDirectionalStreaming(c BiDirectionalAPI.GreetServiceClient) {

	fmt.Println("Starting up client streaming")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling the stream GreetManyTimes %v\n", err)
	}

	requests := []*BiDirectionalAPI.GreetEveryoneRequest{
		{
			Greeting: &BiDirectionalAPI.GreetingBiDirectional{
				FirstName: "HiteshKumar",
			},
		},
		{
			Greeting: &BiDirectionalAPI.GreetingBiDirectional{
				FirstName: "ShilajitRoy",
			},
		},
		{
			Greeting: &BiDirectionalAPI.GreetingBiDirectional{
				FirstName: "pakshalBhandari",
			},
		},
		{
			Greeting: &BiDirectionalAPI.GreetingBiDirectional{
				FirstName: "Balegar",
			},
		},
	}

	waitc := make(chan struct{})

	// Sending data

	go func() {
		for _, req := range requests {
			fmt.Printf("sending data now with value %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// Receiving data
	go func() {

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Coudnt read data while recieveing %v\n", err)
				break
			}
			fmt.Println("Data recieved from the server\n", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc

}
