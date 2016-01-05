package main

import "fmt"

func main() {
	prices := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%T", prices)
	prices = append(prices, 7)
	fmt.Println(prices)
}
