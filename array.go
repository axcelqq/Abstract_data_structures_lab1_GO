package main

import "fmt"

type Array struct {
	data []int
}

func (a *Array) Push(val int) {
	a.data = append(a.data, val)
}

func (a *Array) PushAt(idx int, val int) {
	if idx < 0 || idx > len(a.data) {
		return
	}
	a.data = append(a.data[:idx], append([]int{val}, a.data[idx:]...)...)
}

func (a *Array) RemoveAt(idx int) {
	if idx < 0 || idx >= len(a.data) {
		return
	}
	a.data = append(a.data[:idx], a.data[idx+1:]...)
}

func (a *Array) Set(idx int, val int) {
	if idx >= 0 && idx < len(a.data) {
		a.data[idx] = val
	}
}

func (a *Array) Get(idx int) (int, bool) {
	if idx >= 0 && idx < len(a.data) {
		return a.data[idx], true
	}
	return 0, false
}

func (a *Array) Print() {
	fmt.Println(a.data)
}