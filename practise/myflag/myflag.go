package main

// 可以通过os.Args实现参数的解析

import (
	"flag"
	"fmt"
)

var p1 int 
var p2 string 

func init() {
	flag.IntVar(&p1, "n", 0, "demo number")
	flag.StringVar(&p2, "s", "blank", "demo string")
}

func main() {
	flag.Parse()
	fmt.Println("number is: ", p1)
	fmt.Println("string is: ", p2)
}