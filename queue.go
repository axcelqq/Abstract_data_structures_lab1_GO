package main

import "fmt"

type Queue struct {
	data []string
}

func (q *Queue) Push(val string) { // добавление в конец
	q.data = append(q.data, val)
}

func (q *Queue) Pop() { // удаление из начала
	if len(q.data) == 0 {
		return
	}
	q.data = q.data[1:]
}

func (q *Queue) Peek() (string, bool) { // просмотр начала
	if len(q.data) == 0 {
		return "", false
	}
	return q.data[0], true
}

func (q *Queue) Print() {
	fmt.Println(q.data)
}
