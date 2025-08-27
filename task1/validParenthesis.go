package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	testCase := []struct {
		input    string
		expected bool
	}{
		{"{[]}", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"((()))", true},
		{"[{()}]", true},
		{"[", false},
		{"}", false},
	}

	for _, stru := range testCase {
		fmt.Printf("%-10s is or not valid: %t \n", stru.input, isValid(stru.input))
	}
}
