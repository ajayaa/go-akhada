package main

import "fmt"

var current_balance int

func withdraw(amount uint16) int {
	if current_balance < int(amount) {
		return 0
	} else {
		current_balance -= int(amount)
		return int(amount)
	}
}

func withdraw_all_money() int {
	var max uint16
	var total_money int
	max = 1<<16 - 1
	for max != 0 {
		current_withdrawn := withdraw(max)
		total_money += current_withdrawn
		if current_withdrawn == 0 {
			max = max / 2
		}
	}
	return total_money
}

func main() {
	current_balance = 1
	fmt.Println(withdraw_all_money())
}
