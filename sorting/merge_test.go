package sorting

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMergeSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := MergeSort(tc.slice)
			if !cmp.Equal(tc.expected, actual) {
				t.Errorf(cmp.Diff(tc.expected, actual))
			}
		})
	}
}

func TestInPlaceMergeSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InPlaceMergeSort(tc.slice)
			actual := tc.slice
			if !cmp.Equal(tc.expected, actual) {
				t.Errorf(cmp.Diff(tc.expected, actual))
			}
		})
	}
}
func TestParallelMergeSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ParallelMergeSort(tc.slice)
			if !cmp.Equal(tc.expected, actual) {
				t.Errorf(cmp.Diff(tc.expected, actual))
			}
		})
	}
}
