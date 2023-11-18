package main

import (
	"context"
	"log"
	"net"

	pb "github.com/meirgenuine/overengineered-calculator/division-service/proto/gen"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDivisionServiceServer
}

func (s *server) Divide(ctx context.Context, in *pb.DivisionRequest) (*pb.DivisionResponse, error) {
	result := in.Num1 / in.Num2
	return &pb.DivisionResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDivisionServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
