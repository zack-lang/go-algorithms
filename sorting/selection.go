package sorting

/*
SelectionSort sorts an integer slice in ascending order using the Selection Sort algorithm.
It iterates over the input slice, selecting the minimum element from the unsorted portion
and placing it at the beginning of the sorted portion.
Time complexity: O(n^2), where n is the length of the slice.
The algorithm compares each element with every other element in the unsorted portion,
resulting in quadratic time complexity
*/
func SelectionSort(slice []int) []int {
	n := len(slice)

	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}

	return slice
}

/*
Bidirectional sorts an integer slice in ascending order using the Bidirectional Sort algorithm.
It iterates over the input slice from both ends, simultaneously moving inward and swapping adjacent elements
if they are out of order. This process continues until the slice is sorted.
Time complexity: O(n^2), where n is the length of the slice.
The algorithm iterates over the slice multiple times, with each iteration potentially requiring
comparisons and swaps, resulting in quadratic time complexity.
*/
func Bidirectional(slice []int) []int {
	left := 0
	right := len(slice) - 1

	for left <= right {
		for i := left; i < right; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
		right--

		for i := right; i > left; i-- {
			if slice[i] < slice[i-1] {
				slice[i], slice[i-1] = slice[i-1], slice[i]
			}
		}
		left++
	}

	return slice
}
