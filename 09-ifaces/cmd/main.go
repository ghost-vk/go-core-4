package main

import (
	"io"
	"os"

	"github.com/go-core-4/09-ifaces/pkg/customer"
	"github.com/go-core-4/09-ifaces/pkg/employee"
)

func main() {
	// c1 := customer.New(12)
	// c2 := customer.New(13)
	// e1 := employee.New(20)
	// max := maxAge(c1, c2, e1)
	// fmt.Printf("max age: %d\n", max)

	// Выведет только "bla bla bla end".
	writeStrings(os.Stdout, "bla", " bla bla", 34, struct{ age int }{age: 34}, " end")
}

type Ager interface {
	Age() int
}

func maxAge(peoples ...Ager) int {
	var max int
	for _, p := range peoples {
		a := p.Age()
		if a > max {
			max = a
		}
	}
	return max
}

// То же, что и в maxAge, но типы данных не имеют никаких методов
// и возвращают сам объект, а не возраст.
func maxAge2(unknown ...any) any {
	idx := -1
	max := 0

	for i, p := range unknown {
		// Особая конструкция для проверки типа в switch.
		switch v := p.(type) {
		case *customer.Customer:
			if v.AgeInt >= max {
				idx = i
				max = v.AgeInt
			}
		case *employee.Employee:
			if v.AgeInt >= max {
				idx = i
				max = v.AgeInt
			}
		}
	}

	if idx == -1 {
		return nil
	}

	return unknown[idx]
}

func writeStrings(w io.Writer, strs ...any) {
	for _, s := range strs {
		str, ok := s.(string)
		if !ok {
			continue
		}
		w.Write([]byte(str))
	}
}
