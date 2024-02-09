package sorting

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInsertionSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := InsertionSort(tc.slice)
			if !cmp.Equal(tc.expected, actual) {
				t.Errorf(cmp.Diff(tc.expected, actual))
			}
		})
	}
}

func TestBinaryInsertionSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := BinaryInsertionSort(tc.slice)
			if !cmp.Equal(tc.expected, actual) {
				t.Errorf(cmp.Diff(tc.expected, actual))
			}
		})
	}
}
