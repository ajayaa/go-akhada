package main

func (l *ll) reverse() {
	first := l.head
	second := l.head.next
	third := l.head.next.next
	for third != nil {
		change_direction(first, second)
		first = second
		second = third
		third = third.next
	}
	change_direction(first, second)
	//temp := l.head
	//l.head = l.tail
	//l.tail = temp
	l.head, l.tail = l.tail, l.head
}

func change_direction(first *node, second *node) {
	second.next = first
	if first.prev == nil {
		first.next = nil
	}
	if second.next == nil {
		second.prev = nil
	}
}

func main() {
	l := ll{}
	for i := 1; i <= 10; i++ {
		l.add(i)
	}
	//l.print()
	l.reverse()
	l.print()
}
