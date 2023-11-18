package main

import (
	"context"
	"log"
	"time"

	pb "github.com/meirgenuine/overengineered-calculator/multiplication-service/proto/gen"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMultiplicationServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Multiply(ctx, &pb.MultiplicationRequest{Num1: 2, Num2: 4})
	if err != nil {
		log.Fatalf("failed to multiply: %v", err)
	}
	log.Printf("result: %v", r.GetResult())
}
