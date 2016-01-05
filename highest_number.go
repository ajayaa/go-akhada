package main

import "fmt"

func main() {
	var a uint64
	a = 1
	for i := 63; i > 0; i-- {
		a = (a << 1) | 1
	}
	fmt.Println(a)
}
