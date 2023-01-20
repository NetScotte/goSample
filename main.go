package main

import (
	"fmt"
	"sync"
)

var onece sync.Once

func Job() {
	fmt.Println("run job")
}

func RunJob() {
	onece.Do(Job)
}

func main() {
	for i := 0; i < 5; i++ {
		RunJob()
	}
}
