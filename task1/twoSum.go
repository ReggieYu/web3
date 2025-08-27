package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := numMap[complement]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}

	return nil
}

func main() {
	testCase := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 3, 4}, 7},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
		{[]int{1, 2, 3, 4}, 10},
	}
	for _, test := range testCase {
		fmt.Printf("before sum %v \n", test)
		res := twoSum(test.nums, test.target)
		fmt.Printf("after sum %v\n", res)
	}
}
