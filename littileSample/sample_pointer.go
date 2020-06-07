package main

import "fmt"

// 可以所有的地方都不用&, *这种特殊符号
// 为了节省内存，避免拷贝，使用指针
type Service struct {
	Name string
	// 存放Node的指针
	Nodes []*Node
}

type Node struct {
	Id string
	Ip string
}

func main() {
	// 为了避免拷贝，这里的node是指针
	node := &Node{
		Id: "0000111",
		Ip: "192.168.1.1",
	}

	service := &Service{
		Name: "test",
		// 这里的node是指针，这里的[]*Node含义同line:10
		Nodes: []*Node{
			node,
		},
	}

	fmt.Printf("%+v\n", service)
	// 这里无论是否
	fmt.Printf("IP: %s", service.Nodes[0].Ip)
}