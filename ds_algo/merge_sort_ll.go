package main

import "fmt"

func merge_sort(l *ll) *ll {
	len := l.length()
	l.head = merge_sort_nodes(l.head, len)
	temp := l.head
	for i := 0; i < len-1; i++ {
		temp = temp.next
	}
	temp.next = nil
	return l
}

func merge_sort_nodes(n *node, len int) *node {
	_ = "breakpoint"
	if len == 1 || len == 0 {
		return n
	}
	first, second := split_in_two(n, len)
	if len%2 == 0 {
		return merge_recurse(merge_sort_nodes(first, len/2), merge_sort_nodes(second, len/2), len/2, len/2)
	}
	return merge_recurse(merge_sort_nodes(first, len/2), merge_sort_nodes(second, len/2+1), len/2, len/2+1)
}

func split_in_two(n *node, len int) (*node, *node) {
	if len == 0 || n == nil {
		return nil, nil
	}
	if len == 1 {
		return n, nil
	}
	first := n
	second := n
	for i := 0; i < len/2; i++ {
		second = second.next
	}
	return first, second
}

func merge_recurse(n1 *node, n2 *node, len1 int, len2 int) *node {
	if len1 == 0 {
		return n2
	}
	if len2 == 0 {
		return n1
	}
	if n1.data < n2.data {
		n1.next = merge_recurse(n1.next, n2, len1-1, len2)
		return n1
	}
	n2.next = merge_recurse(n1, n2.next, len1, len2-1)
	return n2
}

func main() {
	fmt.Println("Hi there!")
	l1 := ll{}
	a := [6]int{2, 8, 3, 1, 7, 6}
	for i := 0; i <= 5; i++ {
		l1.add(a[i])
	}
	merge_sort(&l1)
	l1.print()
}
