package main

// switch case示例

import (
	"fmt"
)

func main() {
	a := "hello"
	var i interface{}
	i = a
	switch i.(type) {
	case int:
		fmt.Println("a is int")
	case float32:
		fmt.Println("a is float32")
	case string:
		fmt.Println("a is string")
	default:
		fmt.Println("unknow a type")
	}
}
