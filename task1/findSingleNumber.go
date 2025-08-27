package main

import "fmt"

func findSingleNumber(nums []int) int {
	fre := make(map[int]int)

	//calculate every num count
	for _, num := range nums {
		fre[num]++
	}

	// find only once number
	for num, count := range fre {
		if count == 1 {
			return num
		}
	}

	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 4, 3, 2}
	unique := findSingleNumber(nums)
	fmt.Printf("the unique once number: %d", unique)
}
