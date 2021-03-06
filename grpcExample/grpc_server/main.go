package main

import (
	"context"
	pb "github.com/netscotte/goSample/grpcExample/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// 实现proto文件中的服务
type server struct {

}


func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (out *pb.HelloReply, err error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// 指定我们想监听的端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建grpc的实例
	s := grpc.NewServer()
	// 在grpc服务器上注册我们的服务实例
	pb.RegisterGreeterServer(s, &server{})
	// 调用
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}