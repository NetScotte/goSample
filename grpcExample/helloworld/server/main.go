package main

import (
	"context"
	"fmt"
	pb "github.com/netscotte/goSample/grpcExample/helloworld/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	address = "localhost:9090"
)

// 实现proto文件中的服务
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (out *pb.HelloReply, err error) {
	log.Printf("Received: %v", in.Name)
	msg := fmt.Sprintf("Hello %v", in.Name)
	log.Printf("Send: %v", msg)
	return &pb.HelloReply{Message: msg}, nil
}

func main() {
	// 指定我们想监听的端口
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("grpc listen %v", lis.Addr())
	// 创建grpc的实例
	s := grpc.NewServer()
	// 在grpc服务器上注册我们的服务实例
	pb.RegisterGreeterServer(s, &server{})
	// 调用
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
