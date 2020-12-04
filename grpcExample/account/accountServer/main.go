package main

import (
	"context"
	pb "github.com/netscotte/goSample/grpcExample/account"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50052"
)

// 实现proto文件中的服务
type server struct {
    *pb.UnimplementedAccountServerServer
}


func (s *server) GetUserById(ctx context.Context, in *pb.Id) (out *pb.UserInfo, err error) {
	log.Printf("client request for user id: %v\n", in.Id)
	return &pb.UserInfo{Name: "netliu", Age: 21, Sex: "男"}, nil
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
	pb.RegisterAccountServerServer(s, &server{})
	// 调用
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}