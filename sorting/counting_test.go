package sorting

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_CountingSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			CountingSort(tc.slice)
			actual := tc.slice
			if !cmp.Equal(tc.expected, actual) {
				t.Errorf(cmp.Diff(tc.expected, actual))
			}
		})
	}
}
