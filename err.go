package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 2
	fmt.Println(m)
	val, err := m[2]
	fmt.Printf("%T\n", err)
	fmt.Println(err)
	fmt.Println(val)
}
