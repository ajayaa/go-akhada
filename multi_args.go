package main

import "fmt"

func sum(nums ...int) (int, int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total, 0
}

func main() {
	fmt.Print("Hi there! Welcome back.\n")
	nums := []int{1, 2, 3, 4}
	total := sum(nums...)
	fmt.Println(total)
}
