package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
)

const Topic = "TopicTest"

var mode string

func init() {
	flag.StringVar(&mode, "m", "", "指定运行方式client/server")
}

func Producer() {
	c, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2),
	)
	// 事务消息
	// rocketmq.NewTransactionProducer()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	data := map[string]interface{}{
		"id":      1,
		"service": "test",
		"status":  0,
	}

	byteData, err := json.Marshal(data)
	msg := primitive.Message{
		Topic: Topic,
		Body:  byteData,
	}
	// 带上tag
	msg.WithTag("test")

	// 延迟消息
	// msg.WithDelayTimeLevel(3)

	// 批量消息, msg为分片即可
	sendResult, err := c.SendSync(context.Background(), &msg)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(sendResult.String())
	}

}

func Consumer() {
	sig := make(chan os.Signal)
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 指定selector, 获取tag
	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "TagA || TagC",
	}
	err = c.Subscribe(Topic, selector, consumerHandler)

	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	<-sig
	err = c.Shutdown()
	if err != nil {
		fmt.Println(err)
	}
}

func consumerHandler(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for i := range msgs {
		fmt.Println(msgs[i])
	}
	return consumer.ConsumeSuccess, nil
}

func main() {
	flag.Parse()
	if mode == "" {
		flag.Usage()
		return
	}
	switch mode {
	case "server":
		Producer()
	case "client":
		Consumer()
	default:
		flag.Usage()
	}
	return
}
