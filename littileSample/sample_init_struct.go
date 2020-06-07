package main

import "fmt"

type Node struct {
	Id string
	IP string
}

func WithId(id string) NodeOption{
	return func(node *Node) {
		node.Id = id
	}
}

func WithIP(ip string) NodeOption{
	return func(node *Node) {
		node.IP = ip
	}
}

type NodeOption func(node *Node)

func InitNode(opt ...NodeOption) (node *Node) {
	node = &Node{}
	for _, opt := range(opt) {
		opt(node)
	}
	return node
}

func main() {
	node := InitNode(WithId("01"), WithIP("192.168.1.1"))
	fmt.Printf("%+v", node)
}