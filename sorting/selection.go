package sorting

/*
"The important thing is not to stop questioning. Curiosity has its own reason for existing." - Albert Einstein
Sorting is fundamental to many operations in computer science. When it comes to sorting algorithms, there many algorithms to choose from.
Selection sort may not boast top-tier efficiency but performance isn't always a top priority either.
Sometimes you just need a solution that is simple, elegant, and easy to implement.
So why should you learn how to implement selection sort in Go?
Two reasons for me:
- Selection sort has its own place in the world of algorithms.
- It is also a good starting point for understanding other sorting algorithms.
*/

/*
SelectionSort sorts an integer slice in ascending order using the Selection Sort algorithm.
It iterates over the input slice, selecting the minimum element from the unsorted portion
and placing it at the beginning of the unsorted portion.
Time complexity: O(n^2), where n is the length of the slice.
The algorithm compares each element with every other element in the unsorted portion,
resulting in quadratic time complexity
*/
func SelectionSort(slice []int) []int {
	n := len(slice)

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		slice[i], slice[min] = slice[min], slice[i]
	}

	return slice
}

/*
BidirectionalSelection sorts an integer slice in ascending order using the Bidirectional Selection Sort algorithm.
It iterates over the input slice from both ends, simultaneously moving inward and swapping adjacent elements
if they are out of order. This process continues until the slice is sorted.
Time complexity: O(n), where n is the length of the slice.
The algorithm iterates over the slice multiple times, with each iteration potentially requiring
comparisons and swaps, resulting in linear time complexity.
*/
func BidirectionalSelection(slice []int) []int {
	left, right := 0, len(slice)-1

	for left <= right {
		swapped := false

		// Forward pass
		for i := left; i < right; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				swapped = true
			}
		}
		// If no swaps occurred, the slice is already sorted.
		if !swapped {
			break
		}

		// Backward pass
		for i := right; i > left; i-- {
			if slice[i] < slice[i-1] {
				slice[i], slice[i-1] = slice[i-1], slice[i]
			}
		}

		right--
		left++
	}

	return slice
}
