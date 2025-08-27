package main

import (
	"fmt"
	"strings"
)

func findLongestCommonPrefix(arr []string) string {
	if len(arr) == 0 {
		return ""
	}

	prefix := arr[0]
	for i := 1; i < len(arr); i++ {
		for !strings.HasPrefix(arr[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

func main() {
	arr := []string{"flower", "flow", "flight"}
	commonPrefix := findLongestCommonPrefix(arr)
	fmt.Printf("found the longest common prefix: %s", commonPrefix)
}
