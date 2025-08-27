package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	//sort intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]
		if current[0] < last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

func main() {
	testCase := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {0, 4}},
		{{1, 4}, {2, 3}},
		{{1, 4}, {5, 6}},
		{},
		{{1, 4}},
	}
	for _, tc := range testCase {
		fmt.Printf("before sort int arr %v \t", tc)
		res := merge(tc)
		fmt.Printf("after sort int arr %v \n", res)
	}
}
