package sort

import (
	"cmp"
	"math/bits"
	"slices"
)

func Sort[T cmp.Ordered](slice []T) {
	slices.Sort(slice)
}

// DescSort sorts a slice in a descending order.
func DescSort[T cmp.Ordered](slice []T) {
	slices.SortFunc(slice, func(a, b T) int {
		return cmp.Compare(b, a)
	})
}

func SortBy[T any](slice []T, less func(a, b T) bool) {
	slices.SortFunc(slice, func(a, b T) int {
		if less(a, b) {
			return -1
		}
		if less(b, a) {
			return 1
		}
		return 0
	})
}

func IsSorted[T cmp.Ordered](slice []T) bool {
	return slices.IsSorted(slice)
}

// NthElement reorders the slice in-place such that:
// 1. slice[n] contains the value that would be there at index 'n' if the slice were fully sorted. And it also doesn't neccessarily sort the slice, it just update the slice such that :
// a. All elements before n are <= slice[n].
// b. All elements after n are >= slice[n].
// Implemented via Introselect (3-Way Partitioning + Depth-Limited Heap Fallback).
// Average Time: O(N)
// Worst-Case Time: O(N log N)
func NthElement[T cmp.Ordered](slice []T, n int) (T, bool) {
	if n < 0 || n >= len(slice) {
		var zero T
		return zero, false
	}

	maxDepth := 2 * (bits.Len(uint(len(slice))) - 1)
	introselect(slice, 0, len(slice)-1, n, maxDepth)

	return slice[n], true
}

func introselect[T cmp.Ordered](slice []T, left, right, n, maxDepth int) {
	for left < right {
		if right-left <= 16 {
			insertionSortRange(slice, left, right)
			return
		}

		if maxDepth == 0 {
			heapSortRange(slice, left, right)
			return
		}
		maxDepth--

		mid := left + (right-left)/2
		if slice[right] < slice[left] {
			slice[left], slice[right] = slice[right], slice[left]
		}
		if slice[mid] < slice[left] {
			slice[mid], slice[left] = slice[left], slice[mid]
		}
		if slice[right] < slice[mid] {
			slice[right], slice[mid] = slice[mid], slice[right]
		}
		pivot := slice[mid]

		lt := left
		gt := right
		i := left

		for i <= gt {
			if slice[i] < pivot {
				slice[lt], slice[i] = slice[i], slice[lt]
				lt++
				i++
			} else if slice[i] > pivot {
				slice[i], slice[gt] = slice[gt], slice[i]
				gt--
			} else {
				i++
			}
		}

		if n < lt {
			right = lt - 1
		} else if n > gt {
			left = gt + 1
		} else {
			return
		}
	}
}

// insertionSortRange sorts slice[left : right+1] in-place
func insertionSortRange[T cmp.Ordered](slice []T, left, right int) {
	for i := left + 1; i <= right; i++ {
		key := slice[i]
		j := i - 1
		for j >= left && slice[j] > key {
			slice[j+1] = slice[j]
			j--;
		}
		slice[j+1] = key
	}
}

func heapSortRange[T cmp.Ordered](slice []T, left, right int) {
	sub := slice[left : right+1]
	m := len(sub)

	for i := m/2 - 1; i >= 0; i-- {
		siftDown(sub, i, m)
	}

	for i := m - 1; i > 0; i-- {
		sub[0], sub[i] = sub[i], sub[0]
		siftDown(sub, 0, i)
	}
}

func siftDown[T cmp.Ordered](slice []T, root, end int) {
	for {
		child := 2*root + 1
		if child >= end {
			break
		}
		if child+1 < end && slice[child] < slice[child+1] {
			child++
		}
		if !(slice[root] < slice[child]) {
			break
		}
		slice[root], slice[child] = slice[child], slice[root]
		root = child
	}
}