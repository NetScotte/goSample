package main

/*
基于链表实现queue
*/

var _ Queue = (*LinkQueue)(nil)

type LinkQueue struct {
	Val  interface{}
	Next *LinkQueue
}

func (q *LinkQueue) Add(v interface{}) {
	tempQ := &LinkQueue{
		Val:  v,
		Next: nil,
	}
	tempQ.Next = q.Next
	q.Next = tempQ
}

func (q *LinkQueue) Pop() interface{} {
	s := q.Next.Val
	q.Next = q.Next.Next
	return s
}

func (q *LinkQueue) Empty() bool {
	return q.Next == nil
}
