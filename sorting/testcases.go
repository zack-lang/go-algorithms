package sorting

type testCase struct {
	name     string
	slice    []int
	expected []int
}

var unsorted = []int{2, 5, 1, 3, 6, 10, 4, 8, 7, 9}
var sorted = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var reversed = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var repeated = []int{5, 5, 5, 2, 2, 4, 4, 2, 2, 4, 1, 1, 3}
var empty = []int{}
var oneElement = []int{1}

var testCases = []testCase{
	{
		name:     "Empty slice",
		slice:    empty,
		expected: empty,
	},
	{
		name:     "Unsorted slice",
		slice:    unsorted,
		expected: sorted,
	},
	{
		name:     "Reversed order slice",
		slice:    reversed,
		expected: sorted,
	},
	{
		name:     "Repeated elements slice",
		slice:    repeated,
		expected: []int{1, 1, 2, 2, 2, 2, 3, 4, 4, 4, 5, 5, 5},
	},
	{
		name:     "Sorted slice",
		slice:    sorted,
		expected: sorted,
	},
	{
		name:     "Single-element slice",
		slice:    oneElement,
		expected: oneElement,
	},
}
