package main

import "fmt"

//import "errors"

func pow(x, y int) int {
	if y < 0 {
		return 0
	} else if y == 0 {
		return 1
	} else if y%2 != 0 {
		return pow(x, y/2) * pow(x, y/2) * x
	} else {
		return pow(x, y/2) * pow(x, y/2)
	}
}

func main() {
	fmt.Println(pow(2, 4))
}
