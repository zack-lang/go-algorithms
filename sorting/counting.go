package sorting

/*
CountingSort sorts an integer slice in ascending order using the Counting Sort algorithm.
It counts the occurrences of each number in the input slice using a bucket array,
then reconstructs the sorted array based on the counts.
Time complexity: O(n+k) where n is number of elements and k is range.
The algorithm iterates over the input slice once to count occurrences and then reconstructs
the sorted array from the bucket counts, resulting in linear time complexity.
*/
func CountingSort(slice []int) {
	if len(slice) == 0 {
		return
	}
	max := findMaxVal(slice)
	buckets := make([]int, max+1)

	// Count the number of occurrences of each element in the slice.
	for _, el := range slice {
		buckets[el]++
	}

	idx := 0
	// Rebuild the sorted slice from the bucket counts
	for value, count := range buckets {
		for i := 0; i < count; i++ {
			slice[idx] = value
			idx++
		}
	}
}

func findMaxVal(slice []int) int {
	max := slice[0]
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return max
}
