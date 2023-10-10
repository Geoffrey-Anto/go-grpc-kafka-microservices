package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/geoffrey-anto/golang-microservice-apis/protos"
)

var (
	port          = flag.Int("port", 50051, "The server port")
	file *os.File = nil
)

type server struct {
	pb.UnimplementedLoggerServer
}

func (s *server) SaveLog(ctx context.Context, in *pb.LogSaveRequest) (*pb.LogSaveRespone, error) {
	fmt.Printf("%+v     %+v     %+v\n", in.Id, in.Log, in.Time)
	if file == nil {
		return &pb.LogSaveRespone{
			Success: false,
		}, fmt.Errorf("No file found")
	}

	_, err := file.WriteString(fmt.Sprintf("%+v     %+v     %+v\n", in.Id, in.Log, in.Time))
	if err != nil {
		return &pb.LogSaveRespone{
			Success: false,
		}, fmt.Errorf("Error in running file")
	}

	return &pb.LogSaveRespone{
		Success: true,
	}, nil
}

func main() {
	flag.Parse()

	f, err := os.Create("./logs/logs.txt")
	file = f

	if err != nil {
		log.Fatalf("Error in creating file")
	}

	defer f.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterLoggerServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
