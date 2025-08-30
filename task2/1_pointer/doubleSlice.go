package main

import "fmt"

func doubleSlice(slice *[]int) {
	for i := 0; i < len(*slice); i++ {
		(*slice)[i] *= 2
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("before double slice: %d\n", nums)
	doubleSlice(&nums)
	fmt.Printf("after double slice: %d", nums)

}
