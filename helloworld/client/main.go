package main

import (
	"context"
	"google.golang.org/grpc"
	"helloworld/greet"
	"log"
	"time"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	for range time.Tick(1 * time.Second) {
		doGreet(conn)
	}
}

func doGreet(conn grpc.ClientConnInterface) {
	c := greet.NewGreetServiceClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	msg, err := c.Greet(ctx, &greet.Message{Body: "hello"})
	if err != nil {
		log.Printf("could not greet: %v\n", err)
		return
	}
	log.Printf("Received: %s", msg.GetBody())
}
