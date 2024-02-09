package sorting

/*
InsertionSort sorts an integer slice in ascending order using the Insertion Sort algorithm.
It iterates over the input slice, starting from the second element, and inserts each element
into its correct position among the already sorted elements to the left.
Time complexity: O(n^2), where n is the length of the slice.
The algorithm iterates over the slice multiple times, with each iteration potentially
requiring comparisons and shifts, resulting in quadratic time complexity.
*/

func InsertionSort(slice []int) []int {
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1
		// Move elements of the slice that are greater than key to the right
		for j >= 0 && slice[j] > key {
			slice[j+1] = slice[j]
			j--
		}
		// Insert unsorted element at the correct position
		slice[j+1] = key
	}
	return slice
}

/*
BinaryInsertionSort sorts an integer slice in ascending order using the Binary Insertion Sort algorithm.
It iterates over the input slice, starting from the second element, and uses binary search to find the correct position
to insert each element among the already sorted elements to the left.
Time complexity: O(n^2), where n is the length of the slice.
The algorithm iterates over the slice multiple times, with each iteration potentially
requiring binary search and shifting operations, resulting in quadratic time complexity.
*/
func BinaryInsertionSort(slice []int) []int {
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		low, high := 0, i-1
		// Find the index where the element should be inserted
		index := BinarySearch(slice, low, high, key)
		if index != -1 {
			// Move elements of the slice that are greater than key to the right
			for j := i - 1; j >= index; j-- {
				slice[j+1] = slice[j]
			}
			// Insert unsorted element at the correct position
			slice[index] = key
		}
	}
	return slice
}

func BinarySearch(slice []int, low, high, key int) int {
	for low <= high {
		mid := low + (high-low)/2
		if slice[mid] == key {
			return mid
		} else if slice[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
