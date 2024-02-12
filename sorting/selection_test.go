package sorting

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_SelectionSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := SelectionSort(tc.slice)
			if !cmp.Equal(tc.expected, actual) {
				t.Errorf(cmp.Diff(tc.expected, actual))
			}
		})
	}
}
func Test_BidirectionalSelection(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := BidirectionalSelection(tc.slice)
			if !cmp.Equal(tc.expected, actual) {
				t.Errorf(cmp.Diff(tc.expected, actual))
			}
		})
	}
}
