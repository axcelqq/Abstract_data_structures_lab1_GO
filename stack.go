package main

import "fmt"

type Stack struct {
	data []string
}

func (s *Stack) Push(val string) { // добавление на вершину
	s.data = append(s.data, val)
}

func (s *Stack) Pop() { // удаление с вершины
	if len(s.data) == 0 {
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) Peek() (string, bool) { // просмотр вершины
	if len(s.data) == 0 {
		return "", false
	}
	return s.data[len(s.data)-1], true
}

func (s *Stack) Print() {
	// выводим от вершины к дну
	for i := len(s.data) - 1; i >= 0; i-- {
		fmt.Printf("%s ", s.data[i])
	}
	fmt.Println()
}
