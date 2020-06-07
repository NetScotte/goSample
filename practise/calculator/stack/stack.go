package stack

import (
	"fmt"
)

// Stack不仅存储数字，也要存储计算符号，所以data用string
type Stack struct {
	data [1024]string
	top int 
}

func (s *Stack) Push(n string) {
	if s.top >= 1024 {
		fmt.Println("error, stack is full[1024]")
		return
	}
	s.data[s.top] = n 
	s.top ++
}

func (s *Stack) Pop() (ret string, err error){
	if s.top <= 0 {
		err = fmt.Errorf("error, stack is empty")
		return 
	}
	s.top--
	ret = s.data[s.top]
	s.data[s.top] = ""
	return
}

func (s *Stack) Top() (ret string, err error){
	if s.top == 0 {
		fmt.Errorf("error, stack is empty")
		return 
	}
	ret = s.data[s.top-1]
	return
}

func (s *Stack) Empty() bool{
	if s.top == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Full() bool{
	if s.top == 1024 {
		return true 
	} else {
		return false
	}
}