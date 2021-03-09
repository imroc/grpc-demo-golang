package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"helloworld/greet"
	"log"
	"net"
)

type server struct {
	greet.UnimplementedGreetServiceServer
}

// Greet implements helloworld.GreetServiceServer
func (s *server) Greet(ctx context.Context, in *greet.Message) (*greet.Message, error) {
	log.Printf("Received: %s", in.GetBody())
	return &greet.Message{Body: "world"}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := server{}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	greet.RegisterGreetServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
