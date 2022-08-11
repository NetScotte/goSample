package mycontext

import (
	"context"
	"fmt"
	"time"
)

// 取消goroute任务
// 进行超时控制
// 传递通用参数

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

func doStuff(ctx context.Context) {
	for {
		//time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Printf("done!\n")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("触发了case time")
		}
	}
}
