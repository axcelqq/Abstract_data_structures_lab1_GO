package main

import (
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	Name string
	Type byte // 'M','F','L','S','Q'
	Arr  Array
	Fwd  SList
	Dbl  DList
	Stk  Stack
	Que  Queue
}

type Registry struct {
	entries []*Entry
}

func (r *Registry) Find(name string) *Entry {
	for _, e := range r.entries {
		if e.Name == name {
			return e
		}
	}
	return nil
}

func (r *Registry) GetOrCreate(name string, typ byte) *Entry {
	e := r.Find(name)
	if e != nil {
		if e.Type != typ {
			fmt.Printf("ERROR: name '%s' used for another type\n", name)
			return nil
		}
		return e
	}
	e = &Entry{Name: name, Type: typ}
	r.entries = append(r.entries, e)
	return e
}

func (r *Registry) Save(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("ERROR: cannot save file")
		return
	}
	defer file.Close()

	for _, e := range r.entries {
		prefix := string(e.Type) + "_"
		line := prefix + e.Name + " : "

		var parts []string
		switch e.Type {
		case 'M':
			parts = append(parts, e.Arr.data...)
		case 'F':
			cur := e.Fwd.Head
			for cur != nil {
				parts = append(parts, cur.Data)
				cur = cur.Next
			}
		case 'L':
			cur := e.Dbl.Head
			for cur != nil {
				parts = append(parts, cur.Data)
				cur = cur.Next
			}
		case 'S':
			// Сохраняем снизу вверх
			parts = append(parts, e.Stk.data...)
		case 'Q':
			parts = append(parts, e.Que.data...)
		}
		line += strings.Join(parts, ",")
		file.WriteString(line + "\n")
	}
}

func (r *Registry) Load(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return // Если файла нет — ничего страшного
	}
	defer file.Close()

	buf := make([]byte, 1)
	var line string
	for {
		n, err := file.Read(buf)
		if n == 0 || err != nil {
			if line != "" {
				r.parseLine(line)
			}
			break
		}
		if buf[0] == '\n' {
			r.parseLine(line)
			line = ""
		} else if buf[0] != '\r' {
			line += string(buf[0])
		}
	}
}

func (r *Registry) parseLine(line string) {
	if line == "" {
		return
	}
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return
	}

	left := strings.TrimSpace(parts[0])
	right := strings.TrimSpace(parts[1])
	if len(left) < 3 || left[1] != '_' {
		return
	}

	typ := left[0]
	name := left[2:]

	e := r.GetOrCreate(name, typ)
	if e == nil {
		return
	}

	if right == "" {
		return
	}
	items := strings.Split(right, ",")
	for _, item := range items {
		val := strings.TrimSpace(item)
		if val == "" {
			continue
		}

		switch typ {
		case 'M':
			e.Arr.Push(val)
		case 'F':
			e.Fwd.Push(val)
		case 'L':
			e.Dbl.Push(val)
		case 'S':
			e.Stk.Push(val)
		case 'Q':
			e.Que.Push(val)
		}
	}
}

func (r *Registry) PrintAll() {
	for _, e := range r.entries {
		fmt.Printf("%c_%s : ", e.Type, e.Name)
		switch e.Type {
		case 'M':
			e.Arr.Print()
		case 'F':
			e.Fwd.Print()
		case 'L':
			e.Dbl.Print()
		case 'S':
			e.Stk.Print()
		case 'Q':
			e.Que.Print()
		}
	}
}
