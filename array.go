package main

import "fmt"

type Array struct {
	data []string
}

func (a *Array) Push(val string) { // добавление в конец
	a.data = append(a.data, val)
}

func (a *Array) PushAt(idx int, val string) { // вставка по индексу
	if idx < 0 || idx > len(a.data) {
		return
	}
	a.data = append(a.data[:idx], append([]string{val}, a.data[idx:]...)...)
}

func (a *Array) RemoveAt(idx int) { //удаление по индексу
	if idx < 0 || idx >= len(a.data) {
		return
	}
	a.data = append(a.data[:idx], a.data[idx+1:]...)
}

func (a *Array) Set(idx int, val string) { // замена по индексу
	if idx >= 0 && idx < len(a.data) {
		a.data[idx] = val
	}
}

func (a *Array) Get(idx int) (string, bool) { // чтение по индексу
	if idx >= 0 && idx < len(a.data) {
		return a.data[idx], true
	}
	return "", false
}

func (a *Array) Print() {
	fmt.Println(a.data)
}
