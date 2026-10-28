package main

import "fmt"

type SNode struct {
	Data int
	Next *SNode
}

type SList struct {
	Head *SNode
	Tail *SNode
	Size int
}

func (l *SList) Push(val int) {
	n := &SNode{Data: val}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
	l.Size++
}

func (l *SList) AddHead(val int) {
	n := &SNode{Data: val, Next: l.Head}
	l.Head = n
	if l.Tail == nil {
		l.Tail = n
	}
	l.Size++
}

func (l *SList) RemoveHead() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	if l.Head == nil {
		l.Tail = nil
	}
	l.Size--
}

func (l *SList) RemoveTail() {
	if l.Head == nil {
		return
	}
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		l.Size--
		return
	}
	cur := l.Head
	for cur.Next != l.Tail {
		cur = cur.Next
	}
	cur.Next = nil
	l.Tail = cur
	l.Size--
}

func (l *SList) RemoveValue(val int) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == val {
		l.RemoveHead()
		return
	}
	cur := l.Head
	for cur.Next != nil && cur.Next.Data != val {
		cur = cur.Next
	}
	if cur.Next != nil {
		if cur.Next == l.Tail {
			l.Tail = cur
		}
		cur.Next = cur.Next.Next
		l.Size--
	}
}

func (l *SList) Search(val int) int {
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

func (l *SList) Print() {
	cur := l.Head
	for cur != nil {
		fmt.Printf("%d ", cur.Data)
		cur = cur.Next
	}
	fmt.Println()
}