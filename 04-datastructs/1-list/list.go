// Реализация двусвязного списка вместе с базовыми операциями.
package list

import (
	"fmt"
)

// List - двусвязный список.
type List struct {
	root *Elem
}

// Elem - элемент списка.
type Elem struct {
	Val        interface{}
	next, prev *Elem
}

// New создаёт список и возвращает указатель на него.
func New() *List {
	var l List
	l.root = &Elem{}
	l.root.next = l.root
	l.root.prev = l.root
	return &l
}

// Push вставляет элемент в начало списка.
func (l *List) Push(e *Elem) *Elem {
	e.prev = l.root
	e.next = l.root.next
	l.root.next = e
	if e.next != l.root {
		e.next.prev = e
	}
	return e
}

// String реализует интерфейс fmt.Stringer представляя список в виде строки.
func (l *List) String() string {
	el := l.root.next
	var s string
	for el != l.root {
		s += fmt.Sprintf("%v ", el.Val)
		el = el.next
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

// Pop удаляет первый элемент списка.
func (l *List) Pop() *Elem {
	todel := l.root.next
	if todel == nil {
		return nil
	}
	if todel.next != nil {
		l.root.next = todel.next
		todel.next.prev = l.root
	} else {
		l.root.next = l.root
	}
	return todel
}

// Reverse разворачивает список.
func (l *List) Reverse() *List {
	if l.root.next == l.root {
		return l
	}

	current := l.root.next

	for {
		if current == l.root {
			break
		}
		current.prev, current.next = current.next, current.prev
		current = current.prev
		if current.next == l.root {
			l.root.next = current
		}
	}

	return l
}
