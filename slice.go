package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println("a is ", a)
	fmt.Printf("a is of type: %T\n", a)
	b := make([]int, 0, 5)
	fmt.Println("b is ", b)
}
