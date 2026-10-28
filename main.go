package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "data.txt"
	reg := &Registry{}
	reg.Load(filename)

	fmt.Println("=== Named structures interface (Go) ===")
	fmt.Println("Type HELP for commands, EXIT to quit.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() { break }
		line := scanner.Text()
		if line == "" { continue }

		words := strings.Fields(line)
		if len(words) == 0 { continue }
		
		cmd := words[0]
		if cmd == "EXIT" || cmd == "exit" {
			reg.Save(filename)
			break
		}

		changed := processCommand(words, reg)
		if changed {
			reg.Save(filename)
		}
	}
}

func processCommand(w []string, reg *Registry) bool {
	cmd := w[0]
	changed := false

	// --- Массив (M) ---
	if cmd == "MPUSH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'M')
		if e != nil { e.Arr.Push(parseInt(w[2])); changed = true }
	} else if cmd == "MPUSHAT" && len(w) >= 4 {
		e := reg.GetOrCreate(w[1], 'M')
		if e != nil { e.Arr.PushAt(parseInt(w[2]), parseInt(w[3])); changed = true }
	} else if cmd == "MDEL" && len(w) >= 3 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'M' { e.Arr.RemoveAt(parseInt(w[2])); changed = true }
	} else if cmd == "MSET" && len(w) >= 4 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'M' { e.Arr.Set(parseInt(w[2]), parseInt(w[3])); changed = true }
	} else if cmd == "MGET" && len(w) >= 3 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'M' {
			if v, ok := e.Arr.Get(parseInt(w[2])); ok { fmt.Println("Value:", v) }
		}
	} else if cmd == "MPRINT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'M' { e.Arr.Print() }

	// --- Односвязный (F) ---
	} else if cmd == "FPUSH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'F')
		if e != nil { e.Fwd.Push(parseInt(w[2])); changed = true }
	} else if cmd == "FADDH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'F')
		if e != nil { e.Fwd.AddHead(parseInt(w[2])); changed = true }
	} else if cmd == "FDEL" && len(w) >= 3 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'F' { e.Fwd.RemoveValue(parseInt(w[2])); changed = true }
	} else if cmd == "FDELH" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'F' { e.Fwd.RemoveHead(); changed = true }
	} else if cmd == "FDELT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'F' { e.Fwd.RemoveTail(); changed = true }
	} else if cmd == "FGET" && len(w) >= 3 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'F' { fmt.Println("Index:", e.Fwd.Search(parseInt(w[2]))) }
	} else if cmd == "FPRINT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'F' { e.Fwd.Print() }

	// --- Двусвязный (L) ---
	} else if cmd == "LPUSH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'L')
		if e != nil { e.Dbl.Push(parseInt(w[2])); changed = true }
	} else if cmd == "LADDH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'L')
		if e != nil { e.Dbl.AddHead(parseInt(w[2])); changed = true }
	} else if cmd == "LDEL" && len(w) >= 3 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'L' { e.Dbl.RemoveValue(parseInt(w[2])); changed = true }
	} else if cmd == "LDELH" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'L' { e.Dbl.RemoveHead(); changed = true }
	} else if cmd == "LDELT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'L' { e.Dbl.RemoveTail(); changed = true }
	} else if cmd == "LPRINT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'L' { e.Dbl.Print() }
	} else if cmd == "LPRINTB" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'L' { e.Dbl.PrintBackward() }

	// --- Стек (S) ---
	} else if cmd == "SPUSH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'S')
		if e != nil { e.Stk.Push(parseInt(w[2])); changed = true }
	} else if cmd == "SPOP" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'S' { e.Stk.Pop(); changed = true }
	} else if cmd == "SPEEK" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'S' {
			if v, ok := e.Stk.Peek(); ok { fmt.Println("Top:", v) }
		}
	} else if cmd == "SPRINT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'S' { e.Stk.Print() }

	// --- Очередь (Q) ---
	} else if cmd == "QPUSH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'Q')
		if e != nil { e.Que.Push(parseInt(w[2])); changed = true }
	} else if cmd == "QPOP" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'Q' { e.Que.Pop(); changed = true }
	} else if cmd == "QPEEK" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'Q' {
			if v, ok := e.Que.Peek(); ok { fmt.Println("Front:", v) }
		}
	} else if cmd == "QPRINT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'Q' { e.Que.Print() }

	// --- Дерево (T) ---
	} else if cmd == "TINSERT" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'T')
		if e != nil { e.Tree.Insert(parseInt(w[2])); changed = true }
	} else if cmd == "TGET" && len(w) >= 3 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'T' {
			if e.Tree.Search(parseInt(w[2])) { fmt.Println("Found") } else { fmt.Println("Not found") }
		}
	} else if cmd == "TINORDER" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'T' { e.Tree.PrintInOrder() }
	} else if cmd == "TPREORDER" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'T' { e.Tree.PrintPreOrder() }
	} else if cmd == "TPOSTORDER" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'T' { e.Tree.PrintPostOrder() }
	} else if cmd == "TLEVEL" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'T' { e.Tree.PrintLevelOrder() }

	// --- Общие ---
	} else if cmd == "PRINT" {
		reg.PrintAll()
	} else if cmd == "HELP" {
		printHelp()
	} else {
		fmt.Println("ERROR: unknown command")
	}

	return changed
}

func parseInt(s string) int {
	var v int
	fmt.Sscanf(s, "%d", &v)
	return v
}

func printHelp() {
	fmt.Println("=== Команды ===")
	fmt.Println("Массив   (M): MPUSH <имя> v | MPUSHAT <имя> i v | MDEL <имя> i | MSET <имя> i v")
	fmt.Println("              MGET <имя> i | MPRINT <имя>")
	fmt.Println("Список   (F): FPUSH <имя> v | FADDH <имя> v | FDEL <имя> v | FDELH <имя> | FDELT <имя>")
	fmt.Println("              FGET <имя> v | FPRINT <имя>")
	fmt.Println("Список   (L): LPUSH <имя> v | LADDH <имя> v | LDEL <имя> v | LDELH <имя> | LDELT <имя>")
	fmt.Println("              LPRINT <имя> | LPRINTB <имя>")
	fmt.Println("Стек     (S): SPUSH <имя> v | SPOP <имя> | SPEEK <имя> | SPRINT <имя>")
	fmt.Println("Очередь  (Q): QPUSH <имя> v | QPOP <имя> | QPEEK <имя> | QPRINT <имя>")
	fmt.Println("Дерево   (T): TINSERT <имя> v | TGET <имя> v")
	fmt.Println("              TINORDER <имя> | TPREORDER <имя> | TPOSTORDER <имя> | TLEVEL <имя>")
	fmt.Println("Общие       : PRINT | HELP | EXIT")
}