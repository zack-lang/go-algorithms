package searching

import (
	"fmt"
	"testing"
)

type testCase struct {
	slice    []int
	target   int
	expected bool
}

func TestBinaryContains(t *testing.T) {
	sortedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testCases := []testCase{
		{sortedSlice, 1, true},
		{sortedSlice, 2, true},
		{sortedSlice, 3, true},
		{sortedSlice, 4, true},
		{sortedSlice, 5, true},
		{sortedSlice, 6, true},
		{sortedSlice, 7, true},
		{sortedSlice, 8, true},
		{sortedSlice, 9, true},
		{sortedSlice, 10, true},
		{sortedSlice, 11, false},
		{sortedSlice, -77, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Slice %v, Target %d", tc.slice, tc.target), func(t *testing.T) {
			got := BinaryContains(tc.slice, tc.target)
			if got != tc.expected {
				t.Errorf("Slice %v with target %d, expected %v but got %v", tc.slice, tc.target, tc.expected, got)
			}
		})
	}
}
