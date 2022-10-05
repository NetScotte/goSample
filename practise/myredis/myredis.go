package myredis

import (
	"context"
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
