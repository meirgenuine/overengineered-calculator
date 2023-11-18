package main

import (
	"context"
	"log"
	"time"

	pb "github.com/meirgenuine/overengineered-calculator/subtraction-service/proto/gen"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSubtractionServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Subtract(ctx, &pb.SubtractRequest{Num1: 9.5, Num2: 4})
	if err != nil {
		log.Fatalf("failed to subtract: %v", err)
	}
	log.Printf("result: %v", r.GetResult())
}
