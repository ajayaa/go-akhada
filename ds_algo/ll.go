package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

func (n *node) str() string {
	return fmt.Sprintf("%d", n.data)
}

type ll struct {
	head *node
	tail *node
}

func (l *ll) add(val int) {
	if l.head == nil {
		n := node{data: val}
		l.head = &n
		l.tail = &n
	} else {
		n := node{data: val}
		l.tail.next = &n
		l.tail = &n
		n.prev = l.tail
	}
}

func (l *ll) print() {
	for temp := l.head; temp != nil; temp = temp.next {
		if temp.next != nil {
			fmt.Print(temp.str(), "->")
		} else {
			fmt.Print(temp.str())
		}
	}
	fmt.Println()
}

func (l *ll) length() int {
	len := 0
	for temp := l.head; temp != nil; temp = temp.next {
		len++
	}
	return len
}
