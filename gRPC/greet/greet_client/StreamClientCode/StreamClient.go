package main

import (
	"context"
	"fmt"
	"gRPC/greet/Stream/GreetAPIManyTimes"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i a client")

	ls, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial to localhost:5000 with error %v", err)
	}

	// makes sure the client connection is closed at the end
	defer ls.Close()

	c := GreetAPIManyTimes.NewGreetServiceClient(ls)

	// fmt.Printf("Created client %f", c)
	doClientStreaming(c)

}

func doClientStreaming(c GreetAPIManyTimes.GreetServiceClient) {

	fmt.Println("Starting up client streaming")

	req := &GreetAPIManyTimes.GreetRequestManyTimes{
		Greeting: &GreetAPIManyTimes.Greeting{
			FirstName:  "Hitesh Kumar",
			SecondName: "Balegar",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling the stream GreetManyTimes %v", err)
	}
	for {

		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream from server %v", err)
		}

		log.Println("Recieved response from the server", msg.GetResult())
	}

}
