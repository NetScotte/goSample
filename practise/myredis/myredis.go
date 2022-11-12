package myredis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"time"
)

func Basic() {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	value, err := rdb.Get(context.Background(), "test_aa").Result()
	if err == nil {
		log.Infof("获取到了值, %v", value)
	} else if err == redis.Nil {
		log.Info("没有获取到, 设置值")
		err = rdb.Set(context.Background(), "test_aa", "haha", time.Second*60).Err()
		if err != nil {
			log.Error(err)
			return
		} else {
			log.Info("设置成功")
		}
	} else {
		log.Errorf("连接出错: %v", err)
	}
}

// SortedSet 有序集合功能测试
func SortedSet() {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	key := "liufy47"
	members := []*redis.Z{
		&redis.Z{
			Score:  12,
			Member: "12",
		},
		&redis.Z{
			Score:  56,
			Member: "56",
		},
		&redis.Z{
			Score:  35,
			Member: "35",
		},
		&redis.Z{
			Score:  5,
			Member: "5",
		},
		&redis.Z{
			Score:  11,
			Member: "11",
		},
		&redis.Z{
			Score:  27,
			Member: "27",
		},
	}
	ctx := context.Background()
	for _, item := range members {
		rdb.ZAdd(ctx, key, item)
	}

	cmd := rdb.ZRevRange(ctx, key, 0, 4)
	result, err := cmd.Result()
	if err != nil {
		fmt.Println(err)
	}
	for index, s := range result {
		fmt.Printf("rank %v, value: %v\n", index, s)
	}

}
