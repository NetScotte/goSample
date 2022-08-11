package main

/*
格式化输入
Scanf
Scanln
Scanf
Scanln
格式化输出
Printf		向stdout中输出内容
Println
Sprintf		构造字符串
Sprintln
Fprintf		向可写入对象中写入数据
Fprintln

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
