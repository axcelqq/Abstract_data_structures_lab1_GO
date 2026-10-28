package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() {
	if len(s.data) == 0 { return }
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) Peek() (int, bool) {
	if len(s.data) == 0 { return 0, false }
	return s.data[len(s.data)-1], true
}

func (s *Stack) Print() {
	// Выводим от вершины к дну
	for i := len(s.data) - 1; i >= 0; i-- {
		fmt.Printf("%d ", s.data[i])
	}
	fmt.Println()
}