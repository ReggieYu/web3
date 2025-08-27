package main

import "fmt"

func plusOne(nums []int) []int {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i]++
			return nums
		}
		nums[i] = 0
	}

	return append([]int{1}, nums...)
}

func main() {
	testCase := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{9, 9, 9},
		{0},
	}
	for _, test := range testCase {
		fmt.Printf("the input arr: %v \t", test)
		resarr := plusOne(test)
		fmt.Printf("the result arr: %v\n", resarr)
	}

}
