package main

import "fmt"

func main() {
	// 赋值
	s1 := "haha"         //	string, 取索引为uint8
	s2 := 'h'            // int32
	s3 := `hh`           // string
	r1 := []rune("haha") //	[]int32, 取索引为int32
	fmt.Printf("type s1: %T, type s2: %T, type s3: %T, type r1: %T\n", s1, s2, s3, r1)
	fmt.Printf("type s1[0]: %T\n", s1[0])
	fmt.Printf("type s1[:1]: %T\n", s1[:1])
	fmt.Printf("type r1[0]: %T\n", r1[0])
	fmt.Printf("string(s1[0]) type is: %T\n", string(s1[0]))

	// 字符串遍历
	var a = "haha"
	count := 0
	for _, value := range a {
		if string(value) == "a" {
			count++
		}
	}
	fmt.Printf("a的数量为: %v\n", count)

	b := []rune("haha")
	count = 0
	for _, value := range b {
		if value == 'a' {
			count++
		}
	}
	fmt.Printf("a的数量为: %v\n", count)

	// 字符与数字
}
