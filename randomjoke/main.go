package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "randomjoke/protos/randomjoke"
)

var port = flag.Int("port", 50052, "The server port")

type server struct {
	pb.UnimplementedRandomJokeServiceServer
}

func (s *server) GetRandomJoke(
	ctx context.Context,
	in *pb.RandomJokeRequest,
) (*pb.RandomJokeResponse, error) {
	for {
		return &pb.RandomJokeResponse{
			Joke: "My sea sickness comes in waves.",
		}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterRandomJokeServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
