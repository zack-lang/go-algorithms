package searching

import "testing"

func TestBinaryContains(t *testing.T) {
	sortedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testCases := []struct {
		slice    []int
		target   int
		expected bool
	}{
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
		{sortedSlice, -77, false},
	}
	for _, tc := range testCases {
		actual := BinaryContains(tc.slice, tc.target)
		if actual != tc.expected {
			t.Errorf("expected %v, got %v", tc.expected, actual)
		}
	}
}
