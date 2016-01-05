package main

import "fmt"

func main() {
	i, j := 1.0000000000001, 1.000000000000
	fmt.Printf("%f\n", i)
	if i < j {
		fmt.Println("Yes i is smaller")
	}
}
