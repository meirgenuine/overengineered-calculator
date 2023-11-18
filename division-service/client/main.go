package main

import (
	"context"
	"log"
	"time"

	pb "github.com/meirgenuine/overengineered-calculator/division-service/proto/gen"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := pb.NewDivisionServiceClient(conn)
	r, err := c.Divide(ctx, &pb.DivisionRequest{Num1: 7.5, Num2: 3})
	if err != nil {
		log.Fatalf("failed to divide: %v", err)
	}
	log.Printf("result: %v", r.GetResult())
}
