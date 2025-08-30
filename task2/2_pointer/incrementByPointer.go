package main

import "fmt"

func incrementby(num *int) {
	*num += 10
}

func main() {
	num := 5
	fmt.Printf("before update num: %d\n", num)
	incrementby(&num)
	fmt.Printf("after update num: %d\n", num)
}
