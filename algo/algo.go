package algo

import (
	"cmp"

	sort "github.com/mohitjoshi-hey/stlgo/algo/sort"
	search "github.com/mohitjoshi-hey/stlgo/algo/search"
)

// All search functions are here
func LowerBound[T cmp.Ordered](slice []T, val T) int {
	return search.LowerBound(slice, val)
}

func UpperBound[T cmp.Ordered](slice []T, val T) int {
	return search.UpperBound(slice, val)
}

func BinarySearch[T cmp.Ordered](slice []T, val T) (int, bool) {
	return search.BinarySearch(slice, val)
}

// All sort functions are here
func Sort[T cmp.Ordered](slice []T) {
	sort.Sort(slice)
}

func DescSort[T cmp.Ordered](slice []T) {
	sort.DescSort(slice)
}

func SortBy[T any](slice []T, less func(a, b T) bool) {
	sort.SortBy(slice, less)
}

func IsSorted[T cmp.Ordered](slice []T) bool {
	return sort.IsSorted(slice)
}

func NthElement[T cmp.Ordered](slice []T, n int) (T, bool) {
	return sort.NthElement(slice, n)
}