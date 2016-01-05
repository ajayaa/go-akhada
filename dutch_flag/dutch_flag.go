/*
DUTCH FLAG PROBLEM with three variables.
This code does sort an array which has only three
distinct values. To show the algorithm here an int
type is used as array values. The array is a slice
here. As Slices are passed by reference in golang
and they are the replacement for arrays in golang.
*/

package dutch_flag

import "fmt"

func print_slice(arr []int) {
	fmt.Println(arr)
}

func modify_slice(arr []int) {
	arr[0] = 5
}

func sort(a []int) {
	low, mid := 0, 0
	high := len(a) - 1
	for mid <= high {
		if a[mid] == 1 {
			mid++
		} else if a[mid] == 0 {
			a[mid], a[low] = a[low], a[mid]
			low++
			mid++
		} else if a[mid] == 2 {
			a[high], a[mid] = a[mid], a[high]
			high--
		} else {
			panic("I don't recognize the content of the array to sort it")
		}
	}
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 1, 2, 0, 0, 0, 1}
	//modify_slice(arr)
	sort(arr)
	fmt.Println(arr)
}
