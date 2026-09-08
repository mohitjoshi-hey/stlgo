package search

import "cmp"

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

// Helpful for finding an element in a sorted array.
func BinarySearch[T cmp.Ordered](slice []T, val T) (int, bool) {
	ind := LowerBound(slice, val)

	if ind < len(slice) && slice[ind] == val {
		return ind, true
	}

	return -1, false
}

// Will be used for finding any element in an unsorted array.
func Find[T comparable](slice []T, val T) (int, bool) {
	for i, v := range slice {
		if v == val {
			return i, true;
		}
	}

	return -1, false;
}

func FindIf[T any](slice[]T, pred func(val T) bool) (int, bool) {
	for i, val := range slice {
		if pred(val) {
			return i, true;
		}
	}

	return -1, false;
}