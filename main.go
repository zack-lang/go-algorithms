package main

import (
	"fmt"

	"github.com/zack-lang/go-algorithms/searching"
)

func main() {
	sortedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	target := 1
	result := searching.BinaryContains(sortedSlice, target)
	if !result {
		fmt.Printf("Target %d not found in %v\n", target, sortedSlice)
	} else {
		fmt.Printf("Target %d found in %v\n", target, sortedSlice)
	}
}
