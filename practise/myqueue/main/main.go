package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
)

/*
练习queue的使用

Put	放入元素
Get	获取元素,

*/

func basic() {
	// 创建一个基本队列
	q := queue.Queue{}
	// 依次放入三个元素
	for i := 0; i < 3; i++ {
		err := q.Put(i)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
	}

	// 依次取出
	for !q.Empty() {
		a, err := q.Get(1)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		fmt.Printf("a=%v\n", a[0])
	}
}

type Node struct {
	Id   int
	Name string
	Sex  string //
	Age  int
}

func (node Node) String() string {
	return fmt.Sprintf("编号: %v, 姓名: %v, 性别: %v, 年龄: %v\n", node.Id, node.Name, node.Sex, node.Age)
}

func (node Node) Error() string {
	return fmt.Sprintf("error")
}

func queueAndStruct() {
	q := queue.Queue{}
	node := Node{11, "net", "男", 22}
	err := q.Put(node)
	if err != nil {
		panic(err)
	}
	items, err := q.Get(1)
	if err != nil {
		panic(err)
	}

	newnode, ok := items[0].(Node)
	if !ok {
		panic(ok)
	}

	fmt.Printf("%v", newnode.Id)
}

func main() {
	queueAndStruct()
}
