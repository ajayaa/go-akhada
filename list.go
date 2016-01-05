package main

import "fmt"

type node struct {
	value int
	next  *node
	prev  *node
}

type ll struct {
	head *node
	tail *node
}

func main() {
	n := node{value: 3}
	fmt.Println(n)
}
