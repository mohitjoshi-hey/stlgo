package algo

import "cmp"

// If all elements are strictly less than val, it returns len(slice).
func LowerBound[T cmp.Ordered](slice []T, val T) int {
	low := 0
	high := len(slice)

	for low < high {
		mid := low + (high-low)/2
		if slice[mid] >= val {
			high = mid;
		} else {
			low = mid + 1;
		}
	}

	return low
}

// If no element is strictly greater than val, it returns len(slice).
func UpperBound[T cmp.Ordered](slice []T, val T) int {
	low := 0
	high := len(slice)

	for low < high {
		mid := low + (high-low)/2
		if slice[mid] > val {
			high = mid;
		} else {
			low = mid + 1;
		}
	}

	return low
}

func BinarySearch[T cmp.Ordered](slice []T, val T) (int, bool) {
	ind := LowerBound(slice, val)

	if ind < len(slice) && slice[ind] == val {
		return ind, true
	}

	return -1, false
}