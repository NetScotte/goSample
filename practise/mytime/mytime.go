package mytime

/*
在进行时间转换时，参考格式用2006-01-02 15:04:05进行指定
 */



import (
	"fmt"
	"time"
)

func Sample_time() {
	// 获取当前时间
	current_time := time.Now()
	fmt.Printf("current time is : %v\n", current_time)
	// 获取当前的时间戳
	current_timestamp := time.Now().Unix()
	fmt.Printf("current timestamp is: %v\n", current_timestamp)
	// 格式化当前时间, 该字符串为该时间的字符串表达形式
	format_current_time := current_time.Format("2006/01/02 15:04:05")
	fmt.Printf("current time is formatted to : %q\n", format_current_time)
	// 获取当前的时区
	fmt.Printf("local timezone is: %s\n", current_time.Location())
	// 获取时间的时间戳
	fmt.Printf("current timestamp is: %v\n", current_time.UTC().Unix())
	// 将时间字符串转化为时间
	time_string := "2019-01-30 15:40:40"
	parser_time, _ := time.Parse("2006-01-02 15:04:05", time_string)
	fmt.Printf("parser %T to %T\n", time_string, parser_time)
	// 将时间戳转化为时间
	timestamp := 1538946597
	parser_timestamp := time.Unix(int64(timestamp), 0)
	fmt.Printf("parser %T(%v) to %T(%s)\n", timestamp, timestamp, parser_timestamp, parser_timestamp)
	// 时间加减
	tenMinutesAfter := time.Now().Add(10*time.Minute)
	fmt.Println("ten minutes after is: ", tenMinutesAfter)
}

func Sample_tiker(limit int) {
	tiker := time.Tick(time.Second)
	for v := range tiker {
		fmt.Println(v)
	}
}