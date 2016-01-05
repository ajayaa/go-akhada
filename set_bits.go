package main

import "fmt"

func main() {
	var a uint64
	a = 18446744073709551615
	fmt.Println(a)
	bits := make([]int, 64)
	i := 63
	for ; a != 0; a = a >> 1 {
		bits[i] = int(a & 1)
		i--
	}
	fmt.Println(bits)
}
