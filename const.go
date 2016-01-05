package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func main() {
	// constants can't be declared with := syntax
	const foo = 3
	fmt.Println(foo)
	fmt.Println(needInt(Small))
}
