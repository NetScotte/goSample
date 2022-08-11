package main

import (
	"context"
	pb "github.com/netscotte/goSample/grpcExample/account"
	"google.golang.org/grpc"
	"log"
	"time"
	xerros "github.com/pkg/errors"
	)

const (
	address = "localhost:50053"
)

func main() {
	// 创建到grpc服务器的channel，实际上并没有发起连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", xerros.Wrapf(err, "failed create channel"))
		return
	}
	defer conn.Close()

	// 创建client
	c := pb.NewAccountServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用服务器的方法
	var id int32 = 111
	log.Printf("request server for user id: %v\n", id)
	r, err := c.GetUserById(ctx, &pb.Id{Id: id})
	if err != nil {
		log.Printf("%v", xerros.Wrapf(err, "failed request server"))
		return
	}
	log.Printf("get response from server: %+v\n", r)
}


