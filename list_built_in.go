package main

import (
	"container/list"
	"fmt"
)

type queue struct {
	l list.List
}

func (q *queue) print() {
	l := q.l
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func (q *queue) push(v interface{}) {
	q.l.PushBack(v)
}

func main() {
	l := list.New()
	l.PushBack(3)
	l.PushBack(5)
	l.PushBack(3)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
