package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
)


func main() {
	var q queue.Queue
	err := q.Put("hello")
	if err != nil {
		fmt.Println(err)
	}
	if ! q.Empty(){
		message, err := q.Get(1)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(message[0])

	}
	fmt.Println(q.Len())
	q.Dispose()
}