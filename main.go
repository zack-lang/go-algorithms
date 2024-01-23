package main

import (
	"fmt"

	"github.com/zack-lang/go-algorithms/searching"
)

func main() {
	sortedSlice := []int{1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 9, 10}

	target := 11

	index := searching.FirstOccurrence(sortedSlice, target)

	if index != -1 {
		fmt.Printf("First occurrence of %d is at index %d\n", target, index)
	} else {
		fmt.Printf("%d is not in the slice.\n", target)
	}
}
