package scheduler

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Name       string
	Delay      int
	CreateTime time.Time
	StartTime  time.Time
	EndTime    time.Time
	Error      error
	Cycle      int
}

type DelayScheduler struct {
	Ring   map[int]*Task
	Index  int
	length int
	lock   sync.RWMutex
}

func NewTask(name string, delay int) *Task {
	return &Task{
		Name:  name,
		Delay: delay,
	}
}

func (t *Task) String() string {
	return fmt.Sprintf("Name: %v, Add Time: %v, Run Time: %v, End Time: %v\n", t.Name, t.CreateTime, t.StartTime,
		t.EndTime)
}

func NewDelayScheduler() *DelayScheduler {
	length := 3600
	return &DelayScheduler{
		Ring:   make(map[int]*Task, length),
		Index:  0,
		length: length,
	}
}

func (s *DelayScheduler) Add(t *Task) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	index := t.Delay%s.length + s.Index
	cycle := t.Delay / s.length
	fmt.Printf("delay: %v, current index: %v, put index: %v, cycle: %v\n", t.Delay, s.Index, index, cycle)
	t.Cycle = cycle
	t.CreateTime = time.Now()
	s.Ring[index] = t
}

func (s *DelayScheduler) Remove(t *Task) {

}

func (s *DelayScheduler) Ticker() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Index += 1
	if s.Index >= s.length {
		s.Index = 0
	}
}

func run(t *Task) {
	defer func() {
		t.EndTime = time.Now()
		if p := recover(); p != nil {
			t.Error = fmt.Errorf("%v", p)
		}
	}()
	t.StartTime = time.Now()
	fmt.Println(t.Name)
}

func (s *DelayScheduler) Run() {
	for {
		select {
		case <-time.Tick(time.Second):
			t := s.Ring[s.Index]
			if t != nil {
				go run(t)
			}
			s.Ticker()
		}
	}
}
