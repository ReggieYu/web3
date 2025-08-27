package main

import (
	"fmt"
)

func isPalimdrone(x int) bool {
	if x < 0 {
		return false
	}

	origin := x
	reversed := 0
	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}

	return reversed == origin
}

func main() {
	testNums := []int{121, -121, 10, 12321, 0}
	for _, num := range testNums {
		fmt.Printf("input number: %d is or not palimdrone: %t\n", num, isPalimdrone(num))
	}
}
