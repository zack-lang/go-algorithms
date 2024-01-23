package searching

import (
	"fmt"
	"testing"
)

type testCase struct {
	slice         []int
	target        int
	expected      bool
	expectedIndex int
}

func TestBinaryContains(t *testing.T) {
	sortedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testCases := []testCase{
		{slice: sortedSlice, target: 1, expected: true},
		{slice: sortedSlice, target: 2, expected: true},
		{slice: sortedSlice, target: 3, expected: true},
		{slice: sortedSlice, target: 4, expected: true},
		{slice: sortedSlice, target: 5, expected: true},
		{slice: sortedSlice, target: 6, expected: true},
		{slice: sortedSlice, target: 7, expected: true},
		{slice: sortedSlice, target: 8, expected: true},
		{slice: sortedSlice, target: 9, expected: true},
		{slice: sortedSlice, target: 10, expected: true},
		{slice: sortedSlice, target: 11, expected: false},
		{slice: sortedSlice, target: -77, expected: false},
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

func TestFirstOccurrence(t *testing.T) {
	slices := [][]int{
		{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10},
	}

	testCases := []testCase{
		{slice: slices[0], target: 1, expectedIndex: 0},
		{slice: slices[1], target: 2, expectedIndex: 1},
		{slice: slices[2], target: 5, expectedIndex: 4},
		{slice: slices[3], target: 6, expectedIndex: 5},
		{slice: slices[4], target: 9, expectedIndex: 8},
		{slice: slices[5], target: 10, expectedIndex: 9},
		{slice: slices[0], target: 11, expectedIndex: -1},
		{slice: slices[0], target: -77, expectedIndex: -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Slice %v, Target %d", tc.slice, tc.target), func(t *testing.T) {
			got := FirstOccurrence(tc.slice, tc.target)
			if got != tc.expectedIndex {
				t.Errorf("Slice %v with target %d, expected %v but got %v", tc.slice, tc.target, tc.expectedIndex, got)
			}
		})
	}
}

func TestLastOccurrence(t *testing.T) {
	slices := [][]int{
		{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10},
	}

	testCases := []testCase{
		{slice: slices[0], target: 1, expectedIndex: 1},
		{slice: slices[1], target: 2, expectedIndex: 2},
		{slice: slices[2], target: 5, expectedIndex: 5},
		{slice: slices[3], target: 6, expectedIndex: 6},
		{slice: slices[4], target: 9, expectedIndex: 9},
		{slice: slices[5], target: 10, expectedIndex: 10},
		{slice: slices[0], target: 11, expectedIndex: -1},
		{slice: slices[0], target: -77, expectedIndex: -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Slice %v, Target %d", tc.slice, tc.target), func(t *testing.T) {
			got := LastOccurrence(tc.slice, tc.target)
			if got != tc.expectedIndex {
				t.Errorf("Slice %v with target %d, expected %v but got %v", tc.slice, tc.target, tc.expectedIndex, got)
			}
		})
	}
}

func TestClosetElement(t *testing.T) {
	sortedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	testCases := []testCase{
		{slice: sortedSlice, target: 0, expectedIndex: 0},
		{slice: sortedSlice, target: 1, expectedIndex: 0},
		{slice: sortedSlice, target: 5, expectedIndex: 4},
		{slice: sortedSlice, target: 6, expectedIndex: 5},
		{slice: sortedSlice, target: 7, expectedIndex: 6},
		{slice: sortedSlice, target: 10, expectedIndex: 9},
		{slice: sortedSlice, target: 11, expectedIndex: 9},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Slice %v, Target %d", tc.slice, tc.target), func(t *testing.T) {
			got := ClosestElement(tc.slice, tc.target)
			if got != tc.expectedIndex {
				t.Errorf("Slice %v with target %d, expected %v but got %v", tc.slice, tc.target, tc.expectedIndex, got)
			}
		})
	}
}
