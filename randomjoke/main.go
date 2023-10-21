package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	pb "server/protos/randomjoke"
	"time"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type server struct {
	pb.UnimplementedRandomJokeServiceServer
}

func (s *server) GetRandomJoke(in *pb.RandomJokeRequest, stream pb.RandomJokeService_GetRandomJokeServer) error {
	for {
		stream.Send(&pb.RandomJokeResponse{Joke: "This is a joke"})
		time.Sleep(time.Duration(in.Timeout) * time.Second)
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
