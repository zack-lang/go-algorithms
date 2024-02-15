package sorting

/*
MergeSort sorts an integer slice in ascending order using the Merge Sort algorithm.
It divides the input slice into two halves, recursively sorts each half, and then merges the sorted halves.
Time complexity: O(n log n), where n is the length of the slice.
The algorithm divides the input slice into halves recursively, resulting in logarithmic time complexity,
and merges the sorted halves in linear time.
*/
func MergeSort(slice []int) []int {
	// Base case:
	// If the slice has 0 or 1 elements, it is already sorted
	if len(slice) <= 1 {
		return slice
	}
	// Divide the slice into two halves
	mid := len(slice) / 2
	left := MergeSort(slice[:mid])
	right := MergeSort(slice[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	// Compare elements from both slices until one is exhausted
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// Append any remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

/*
	InPlaceMergeSort sorts an integer slice in ascending order using the in-place Merge Sort algorithm.
	Time complexity: O(n log n), where n is the length of the slice.
	The algorithm recursively divides the input slice into halves and merges them in place, resulting in logarithmic time complexity.
*/

func InPlaceMergeSort(slice []int) {
	// Base case
	if len(slice) <= 1 {
		return // If the slice has 0 or 1 elements, it is already sorted
	}
	mid := len(slice) / 2
	InPlaceMergeSort(slice[:mid])
	InPlaceMergeSort(slice[mid:])
	mergeInPlace(slice, 0, mid, len(slice))
}

func mergeInPlace(slice []int, left, mid, right int) {
	temp := make([]int, right-left)
	copy(temp, slice[left:right])

	i, j := 0, mid-left
	for k := left; k < right; k++ {
		if i == mid-left { // If the left slice is exhausted
			slice[k] = temp[j]
			j++
		} else if j == right-left || temp[i] <= temp[j] { // If the right slice is exhausted or the current element in the left slice is less than the current element in the right slice
			slice[k] = temp[i]
			i++
		} else {
			slice[k] = temp[j] // If the current element in the left slice is greater than the current element in the right slice
			j++
		}
	}
}
