package searching

/*
Searches for a target value in a sorted slice of integers using binary search. Returns true and nil if found, false and an error if not found.
Time complexity: O(log n), where n is the length of the slice. The algorithm repeatedly divides the search space in half,
resulting in logarithmic search time on average. In the worst case when the target is not present or at the extremes
of the slice, the complexity approaches O(n).
*/
func BinaryContains(slice []int, target int) bool {
	low, high := 0, len(slice)-1

	for low <= high {
		mid := low + (high-low)/2
		if slice[mid] == target {
			return true
		} else if slice[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
