package main

import (
	"context"
	"log"
	"net"

	pb "github.com/meirgenuine/overengineered-calculator/addition-service/proto/gen"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAdditionServiceServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	result := in.Num1 + in.Num2
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdditionServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
