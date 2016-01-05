package main

import "fmt"

func is_power_of_two(num int) bool {
	seen_one := false
	for num != 0 {
		if !seen_one {
			if num%2 != 0 {
				seen_one = true
			}
		} else {
			return false
		}
		num = num >> 1
	}
	return true
}

func main() {
	a := 0
	fmt.Println(is_power_of_two(a))
}
