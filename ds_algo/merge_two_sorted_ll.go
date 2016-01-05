package main

import "fmt"

func merge(l1 *ll, l2 *ll) *ll {
	result := ll{}
	result.head = merge_recurse(l1.head, l2.head)
	return &result
}

func merge_recurse(n1 *node, n2 *node) *node {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.data < n2.data {
		n1.next = merge_recurse(n1.next, n2)
		return n1
	}
	n2.next = merge_recurse(n1, n2.next)
	return n2
}

func main() {
	fmt.Println("Hi there!")
	l1 := ll{}
	l2 := ll{}
	a := [5]int{2, 6, 8, 15, 20}
	b := [5]int{1, 4, 7, 25, 30}
	for i := 0; i <= 4; i++ {
		l1.add(a[i])
		l2.add(b[i])
	}
	merge(&l1, &l2).print()
}
