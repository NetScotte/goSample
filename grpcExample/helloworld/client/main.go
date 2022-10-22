package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	pb "github.com/netscotte/goSample/grpcExample/helloworld/helloworld"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9090"
)

func main() {
	// 创建到grpc服务器的channel
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建client
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用服务器的方法
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "grpc"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("receive Server Greeting: %v", r.Message)
}
