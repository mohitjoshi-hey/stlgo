package algo

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

func BinarySearch[T cmp.Ordered](slice []T, val T) (int, bool) {
	ind := LowerBound(slice, val)

	if ind < len(slice) && slice[ind] == val {
		return ind, true
	}

	return -1, false
}