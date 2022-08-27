package main

/*
主要分为两类方法，一个scan，一个print
前缀包括F和S，后缀包括f和ln
如果前缀是"F", 则指定了io.Writer
如果前缀是"S", 则是输出到字符串
如果后缀是"f", 则指定了format
如果后缀是"ln", 则有换行符
*/

import (
	"fmt"
	"os"
)

func test1() {
	fmt.Println("Sample input-----multi value in one line")
	var a int
	var b string
	var c float32
	fmt.Println("input a b c")
	fmt.Scanf("%d%s%f", &a, &b, &c)
	fmt.Printf("a=%d, b=%s, c=%f\n", a, b, c)
}

func test2() {
	fmt.Println("Sample input-----one line one value")
	var a int
	var b string
	var c float32
	fmt.Println("input a")
	fmt.Scanf("%d\n", &a)
	fmt.Println("input b")
	fmt.Scanf("\n%s\n", &b)
	fmt.Println("input c")
	fmt.Scanf("\n%f\n", &c)
	fmt.Printf("a=%d, b=%s, c=%f\n", a, b, c)
}

func test3() {
	fmt.Println("Sample input-----one line one value")
	var a int
	var b string
	var c float32
	fmt.Println("input ")
	fmt.Scanln(&a)
	fmt.Println("input b")
	fmt.Scanln(&b)
	fmt.Println("input c")
	fmt.Scanln(&c)
	fmt.Printf("a=%d, b=%s, c=%f\n", a, b, c)
}

func os_fmt() {
	var b [16]byte
	os.Stdout.Write([]byte("input some\n"))
	os.Stdin.Read(b[:])
	os.Stdout.Write(b[:])
}

func main() {
	//test1()
	//test2()
	test3()
	//os_fmt()
}
