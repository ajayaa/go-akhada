package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(" I am gonna generate a random number between 1 and 10")
	fmt.Println("it is :", rand.Intn(10))
}
