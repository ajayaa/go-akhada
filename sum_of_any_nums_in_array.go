package main

import "fmt"

func is_sum_present(arr []int, sum int, index int) bool {
	if sum == 0 {
		return true
	} else if sum < 0 || index > len(arr)-1 {
		return false
	}
	return is_sum_present(arr, sum-arr[index], index+1) || is_sum_present(arr, sum, index+1)
}
func main() {
	arr := []int{2, 4, 1, 6, 9, 10}
	fmt.Println(is_sum_present(arr, 13, 0))
}
