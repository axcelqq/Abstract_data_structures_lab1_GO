package main

import "fmt"

type DNode struct {
	Data string
	Prev *DNode
	Next *DNode
}

type DList struct {
	Head *DNode
	Tail *DNode
	Size int
}

func (l *DList) Push(val string) { // добавление в конец
	n := &DNode{Data: val, Prev: l.Tail}
	if l.Tail != nil {
		l.Tail.Next = n
	} else {
		l.Head = n
	}
	l.Tail = n
	l.Size++
}

func (l *DList) AddHead(val string) { // добавление в начало
	n := &DNode{Data: val, Next: l.Head}
	if l.Head != nil {
		l.Head.Prev = n
	} else {
		l.Tail = n
	}
	l.Head = n
	l.Size++
}

func (l *DList) RemoveHead() { // удаление из начала
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	if l.Head != nil {
		l.Head.Prev = nil
	} else {
		l.Tail = nil
	}
	l.Size--
}

func (l *DList) RemoveTail() { // удаление из конца
	if l.Tail == nil {
		return
	}
	l.Tail = l.Tail.Prev
	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil
	}
	l.Size--
}

func (l *DList) RemoveValue(val string) { // удаление первого вхождения
	cur := l.Head
	for cur != nil && cur.Data != val {
		cur = cur.Next
	}
	if cur == nil {
		return
	}

	if cur.Prev != nil {
		cur.Prev.Next = cur.Next
	} else {
		l.Head = cur.Next
	}
	if cur.Next != nil {
		cur.Next.Prev = cur.Prev
	} else {
		l.Tail = cur.Prev
	}
	l.Size--
}

func (l *DList) Search(val string) int { // поиск по значению
	cur := l.Head
	idx := 0
	for cur != nil {
		if cur.Data == val {
			return idx
		}
		cur = cur.Next
		idx++
	}
	return -1
}

func (l *DList) Print() {
	cur := l.Head
	for cur != nil {
		fmt.Printf("%s ", cur.Data)
		cur = cur.Next
	}
	fmt.Println()
}

func (l *DList) PrintBackward() { // вывод в обратном порядке
	cur := l.Tail
	for cur != nil {
		fmt.Printf("%s ", cur.Data)
		cur = cur.Prev
	}
	fmt.Println()
}
