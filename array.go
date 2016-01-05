package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a)
	a[1] = 3
	fmt.Println(a)
}
