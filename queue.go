package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) Push(val int) {
	q.data = append(q.data, val)
}

func (q *Queue) Pop() {
	if len(q.data) == 0 { return }
	q.data = q.data[1:]
}

func (q *Queue) Peek() (int, bool) {
	if len(q.data) == 0 { return 0, false }
	return q.data[0], true
}

func (q *Queue) Print() {
	fmt.Println(q.data)
}