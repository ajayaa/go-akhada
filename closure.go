package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	i := intSeq()
	for j := 0; j < 10; j++ {
		fmt.Println(i())
	}
}
