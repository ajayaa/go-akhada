package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	fmt.Println("x is: ", x)
	var z = 3.0
	var iter = 10
	for i := 0; i < iter; i++ {
		z = z - ((z*z - x) / (2 * z))
		fmt.Println("Now z is: ", z)
	}
	return z
}

func main() {
	var foo = Sqrt(25.0)
	fmt.Println("sqrt is: ", foo, "And the diff is: ", math.Sqrt(25)-foo)
}
