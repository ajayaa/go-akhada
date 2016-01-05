package main

import "fmt"

func verify_sorted_array(arr []int) bool {
	len := len(arr)
	for i := 1; i < len; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 4}
	fmt.Println(verify_sorted_array(arr))
}
