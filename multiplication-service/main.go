package main

import (
	"context"
	"log"
	"net"

	pb "github.com/meirgenuine/overengineered-calculator/multiplication-service/proto/gen"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMultiplicationServiceServer
}

func (s *server) Multiply(ctx context.Context, in *pb.MultiplicationRequest) (*pb.MultiplicationResponse, error) {
	result := in.Num1 * in.Num2
	return &pb.MultiplicationResponse{Result: result}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMultiplicationServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
