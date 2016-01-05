package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	if x := 2; sum == 45 {
		fmt.Println("Yes I am 45 and x is", x)
	}
}
