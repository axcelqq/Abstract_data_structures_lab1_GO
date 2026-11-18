package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Точка входа (поддерживает и интерактивный режим, и флаги --file/--query)
func main() {
	filename := "data.txt"
	query := ""

	// Разбор аргументов командной строки
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "--file" && i+1 < len(args) {
			filename = args[i+1]
			i++
		} else if args[i] == "--query" && i+1 < len(args) {
			query = args[i+1]
			i++
		}
	}

	reg := &Registry{}
	reg.Load(filename)

	// РЕЖИМ 1: Одиночная команда (через --query)
	if query != "" {
		words := strings.Fields(query)
		if len(words) > 0 {
			changed := processCommand(words, reg)
			if changed {
				reg.Save(filename) // автосохранение, если данные изменились
			}
		}
		return // выход из программы после выполнения одной команды
	}

	// РЕЖИМ 2: Интерактивная консоль (если нет флагов)
	fmt.Println("=== Named structures interface (Go) ===")
	fmt.Println("Type HELP for commands, EXIT to quit.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "" {
			continue
		}

		words := strings.Fields(line)
		if len(words) == 0 {
			continue
		}

		cmd := words[0]
		if cmd == "EXIT" || cmd == "exit" {
			reg.Save(filename)
			fmt.Println("Данные сохранены. Выход.")
			break
		}

		changed := processCommand(words, reg)
		if changed {
			reg.Save(filename)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
	}
}

// Обработчик команд
func processCommand(w []string, reg *Registry) bool {
	cmd := w[0]
	changed := false

	// Массив (M)
	if cmd == "MPUSH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'M')
		if e != nil {
			e.Arr.Push(w[2])
			changed = true
		}
	} else if cmd == "MPUSHAT" && len(w) >= 4 {
		e := reg.GetOrCreate(w[1], 'M')
		if e != nil {
			e.Arr.PushAt(parseInt(w[2]), w[3])
			changed = true
		}
	} else if cmd == "MDEL" && len(w) >= 3 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'M' {
			e.Arr.RemoveAt(parseInt(w[2]))
			changed = true
		}
	} else if cmd == "MSET" && len(w) >= 4 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'M' {
			e.Arr.Set(parseInt(w[2]), w[3])
			changed = true
		}
	} else if cmd == "MGET" && len(w) >= 3 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'M' {
			if v, ok := e.Arr.Get(parseInt(w[2])); ok {
				fmt.Println("Value:", v)
			} else {
				fmt.Println("ERROR: index out of range")
			}
		}
	} else if cmd == "MPRINT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'M' {
			e.Arr.Print()
		}

	} else if cmd == "MLEN" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'M' {
			fmt.Println("len =", len(e.Arr.data))
		}

		// Односвязный (F)
	} else if cmd == "FPUSH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'F')
		if e != nil {
			e.Fwd.Push(w[2])
			changed = true
		}
	} else if cmd == "FADDH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'F')
		if e != nil {
			e.Fwd.AddHead(w[2])
			changed = true
		}
	} else if cmd == "FDEL" && len(w) >= 3 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'F' {
			e.Fwd.RemoveValue(w[2])
			changed = true
		}
	} else if cmd == "FDELH" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'F' {
			e.Fwd.RemoveHead()
			changed = true
		}
	} else if cmd == "FDELT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'F' {
			e.Fwd.RemoveTail()
			changed = true
		}
	} else if cmd == "FGET" && len(w) >= 3 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'F' {
			fmt.Println("Index:", e.Fwd.Search(w[2]))
		}
	} else if cmd == "FPRINT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'F' {
			e.Fwd.Print()
		}
	} else if cmd == "FLEN" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'F' {
			fmt.Println("len =", e.Fwd.Size)
		}

		// Двусвязный (L)
	} else if cmd == "LPUSH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'L')
		if e != nil {
			e.Dbl.Push(w[2])
			changed = true
		}
	} else if cmd == "LADDH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'L')
		if e != nil {
			e.Dbl.AddHead(w[2])
			changed = true
		}
	} else if cmd == "LDEL" && len(w) >= 3 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'L' {
			e.Dbl.RemoveValue(w[2])
			changed = true
		}
	} else if cmd == "LDELH" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'L' {
			e.Dbl.RemoveHead()
			changed = true
		}
	} else if cmd == "LDELT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'L' {
			e.Dbl.RemoveTail()
			changed = true
		}
	} else if cmd == "LGET" && len(w) >= 3 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'L' {
			fmt.Println("Index:", e.Dbl.Search(w[2]))
		}
	} else if cmd == "LPRINT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'L' {
			e.Dbl.Print()
		}
	} else if cmd == "LPRINTB" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'L' {
			e.Dbl.PrintBackward()
		}
	} else if cmd == "LLEN" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'L' {
			fmt.Println("len =", e.Dbl.Size)
		}

		// Стек (S)
	} else if cmd == "SPUSH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'S')
		if e != nil {
			e.Stk.Push(w[2])
			changed = true
		}
	} else if cmd == "SPOP" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'S' {
			e.Stk.Pop()
			changed = true
		}
	} else if cmd == "SPEEK" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'S' {
			if v, ok := e.Stk.Peek(); ok {
				fmt.Println("Top:", v)
			} else {
				fmt.Println("ERROR: stack is empty")
			}
		}
	} else if cmd == "SPRINT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'S' {
			e.Stk.Print()
		}
	} else if cmd == "SLEN" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'S' {
			fmt.Println("len =", len(e.Stk.data))
		}

		// Очередь (Q)
	} else if cmd == "QPUSH" && len(w) >= 3 {
		e := reg.GetOrCreate(w[1], 'Q')
		if e != nil {
			e.Que.Push(w[2])
			changed = true
		}
	} else if cmd == "QPOP" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'Q' {
			e.Que.Pop()
			changed = true
		}
	} else if cmd == "QPEEK" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'Q' {
			if v, ok := e.Que.Peek(); ok {
				fmt.Println("Front:", v)
			} else {
				fmt.Println("ERROR: queue is empty")
			}
		}
	} else if cmd == "QPRINT" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'Q' {
			e.Que.Print()
		}
	} else if cmd == "QLEN" && len(w) >= 2 {
		if e := reg.Find(w[1]); e != nil && e.Type == 'Q' {
			fmt.Println("len =", len(e.Que.data))
		}

		// Общие
	} else if cmd == "PRINT" {
		reg.PrintAll()
	} else if cmd == "HELP" {
		printHelp()
	} else {
		fmt.Println("ERROR: unknown command")
	}

	return changed
}

// Вспомогательная функция для перевода строки в число (для индексов)
func parseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

// Вывод справки
func printHelp() {
	fmt.Println("=== Команды ===")
	fmt.Println("Массив   (M): MPUSH <имя> v | MPUSHAT <имя> i v | MDEL <имя> i | MSET <имя> i v")
	fmt.Println("              MGET <имя> i | MPRINT <имя>")
	fmt.Println("Список   (F): FPUSH <имя> v | FADDH <имя> v | FDEL <имя> v | FDELH <имя> | FDELT <имя>")
	fmt.Println("              FGET <имя> v | FPRINT <имя>")
	fmt.Println("Список   (L): LPUSH <имя> v | LADDH <имя> v | LDEL <имя> v | LDELH <имя> | LDELT <имя>")
	fmt.Println("              LGET <имя> v | LPRINT <имя> | LPRINTB <имя>")
	fmt.Println("Стек     (S): SPUSH <имя> v | SPOP <имя> | SPEEK <имя> | SPRINT <имя>")
	fmt.Println("Очередь  (Q): QPUSH <имя> v | QPOP <имя> | QPEEK <имя> | QPRINT <имя>")
	fmt.Println("Общие       : PRINT | HELP | EXIT")
}
