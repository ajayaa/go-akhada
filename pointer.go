package main

import "fmt"

func value_do_not_change(i int) {
	i = 0
}

func value_does_change(i *int) {
	*i = 0
}

func main() {
	i := 42
	fmt.Println(i)
	value_do_not_change(i)
	fmt.Println(i)
	value_does_change(&i)
	fmt.Println(i)
	fmt.Println(&i)
	//fmt.Println("p is: ", p)
	//*p = 21
	//fmt.Println("i is: ", i, j)
}
