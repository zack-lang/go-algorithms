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

/*
FirstOccurrence searches for the leftmost occurrence of a target value in a sorted slice of integers using binary search.
If the target is found, it returns the index of the leftmost occurrence; otherwise, it returns -1.
Time complexity: O(log n), where n is the length of the slice. The algorithm repeatedly divides the search space in half,
resulting in logarithmic search time on average. In the worst case when the target is not present or at the extremes
of the slice, the complexity approaches O(n).
*/
func FirstOccurrence(slice []int, target int) int {
	left, right := 0, len(slice)-1
	occurrence := -1
	for left <= right {
		mid := left + (right-left)/2
		if slice[mid] == target {
			occurrence = mid
			right = mid - 1
		} else if slice[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return occurrence
}

/*
LastOccurrence searches for the rightmost occurrence of a target value in a sorted slice of integers using binary search.
If the target is found, it returns the index of the rightmost occurrence; otherwise, it returns -1.
Time complexity: O(log n), where n is the length of the slice. The algorithm repeatedly divides the search space in half,
resulting in logarithmic search time on average. In the worst case when the target is not present or at the extremes
of the slice, the complexity approaches O(n).
*/
func LastOccurrence(slice []int, target int) int {
	left, right := 0, len(slice)-1
	occ := -1
	for left <= right {
		mid := left + (right-left)/2
		if slice[mid] == target {
			occ = mid
			left = mid + 1
		} else if slice[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return occ
}

/*
ClosestElement finds the index of the element in a sorted slice of integers that is closest to the given target value using binary search.
If the target is present in the slice, it returns the index of the target. Otherwise, it returns the index of the element closest to the target.
Time complexity: O(log n), where n is the length of the slice. The algorithm repeatedly divides the search space in half,
resulting in logarithmic search time on average. In the worst case when the target is not present or at the extremes
of the slice, the complexity approaches O(n).
*/
func ClosestElement(slice []int, target int) int {
	left, right := 0, len(slice)-1
	closestIndex := 0

	for left <= right {
		mid := left + (right-left)/2
		if slice[mid] == target {
			return mid
		}

		diff := abs(slice[mid] - target)

		if closestIndex == -1 || diff < abs(slice[closestIndex]-target) {
			closestIndex = mid
		}

		if slice[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return closestIndex
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
