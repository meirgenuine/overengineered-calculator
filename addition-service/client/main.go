package main

import (
	"log"
	"time"

	pb "github.com/meirgenuine/overengineered-calculator/addition-service/proto/gen"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAdditionServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Add(ctx, &pb.AddRequest{Num1: 4.5, Num2: 5})
	if err != nil {
		log.Fatalf("failed to add: %v", err)
	}
	log.Printf("result: %v", r.GetResult())
}
