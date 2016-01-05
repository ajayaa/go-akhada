package main

import "fmt"
import "sort"

type int_array []int

func sum_of_two_nums(arr []int, sum int) bool {
	sort.Sort(int_array(arr))
	j := len(arr) - 1
	i := 0
	for i < j {
		if arr[i]+arr[j] == sum {
			return true
		} else if arr[i]+arr[j] < sum {
			i++
		} else if arr[i]+arr[j] > sum {
			j--
		}
	}
	return false
}

func (arr int_array) Len() int           { return len(arr) }
func (arr int_array) Swap(i, j int)      { arr[i], arr[j] = arr[j], arr[i] }
func (arr int_array) Less(i, j int) bool { return arr[i] < arr[j] }

func main() {
	arr := []int{4, 8, 3, 5, 12, 18, 7}
	fmt.Println(sum_of_two_nums(arr, 14))
}
