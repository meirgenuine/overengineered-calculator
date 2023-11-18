package main

import (
	"context"
	"log"
	"net"

	pb "github.com/meirgenuine/overengineered-calculator/subtraction-service/proto/gen"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSubtractionServiceServer
}

func (s *server) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	result := in.Num1 - in.Num2
	return &pb.SubtractResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSubtractionServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
