package main

// 命令行参数解析
import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	// you can receive and return a var like this
	// var name = flag.String("name", "everyone", "the greeting object.")
	// 另外还有方法: Int, Bool
}

// 执行示例: go run sample_flag.go -name netliu
func main() {
	flag.Parse()
	fmt.Printf("Hello, %s\n", name)
}