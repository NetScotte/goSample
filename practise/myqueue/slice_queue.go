package main

/*
基于slice实现queue
*/

var _ Queue = (*SliceQueue)(nil)

type SliceQueue struct {
	s   []interface{} // 数据的存储容器
	top int           // 队列中的元素个数
	len int           // slice实际存储的元素个数
}

// Add 入队
func (q *SliceQueue) Add(v interface{}) {
	if q.top >= q.len {
		q.s = append(q.s, v)
		q.len += 1
	} else {
		q.s[q.top] = v
	}
	q.top += 1
}

// Pop 出队，如果为空Pop，则panic
func (q *SliceQueue) Pop() interface{} {
	if q.Empty() {
		panic("队列为空，无法出队")
	} else {
		q.top -= 1
		return q.s[q.top]
	}
}

func (q *SliceQueue) Empty() bool {
	return q.top == 0
}
