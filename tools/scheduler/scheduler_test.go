package scheduler

import (
	"fmt"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	s := NewDelayScheduler()
	go s.Run()

	task1 := NewTask("test", 2)
	task2 := NewTask("wait", 5)
	s.Add(task1)
	s.Add(task2)
	time.Sleep(10 * time.Second)
	fmt.Println(task1)
	fmt.Println(task2)
}
