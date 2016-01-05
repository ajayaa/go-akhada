package main

import "fmt"

/* Assuming the array only contains
positive numbers we return -1 if
the length is less than 2 */

func second_max(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return -1
	}
	var max, next_max int
	max = arr[0]
	next_max = arr[1]
	if arr[0] < arr[1] {
		max, next_max = arr[1], arr[0]
	}
	for i := 2; i < length; i++ {
		if arr[i] > max {
			next_max, max = max, arr[i]
		} else {
			if arr[i] > next_max {
				next_max = arr[i]
			}
		}
	}
	return next_max
}

func main() {
	arr := []int{1, 1, 1}
	fmt.Println(second_max(arr))
}
