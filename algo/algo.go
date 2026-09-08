package algo

import (
	"cmp"

	math "github.com/mohitjoshi-hey/stlgo/algo/math"
	numeric "github.com/mohitjoshi-hey/stlgo/algo/numeric"
	search "github.com/mohitjoshi-hey/stlgo/algo/search"
	sort "github.com/mohitjoshi-hey/stlgo/algo/sort"
)

// Search functions
func LowerBound[T cmp.Ordered](slice []T, val T) int {
	return search.LowerBound(slice, val)
}

func UpperBound[T cmp.Ordered](slice []T, val T) int {
	return search.UpperBound(slice, val)
}

func BinarySearch[T cmp.Ordered](slice []T, val T) (int, bool) {
	return search.BinarySearch(slice, val)
}

func Find[T comparable](slice []T, val T) (int, bool) {
	return search.Find(slice, val)
}

func FindIf[T any](slice []T, pred func(val T) bool) (int, bool) {
	return search.FindIf(slice, pred)
}

// Sort functions
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

// Math functions
func GCD[T math.Integer](a, b T, rest ...T) T {
	return math.GCD(a, b, rest...)
}

func LCM[T math.Integer](a, b T, rest ...T) T {
	return math.LCM(a, b, rest...)
}

func IsPrime[T math.Integer](n T) bool {
	return math.IsPrime(n)
}

func IsEven[T math.Integer](val T) bool {
	return math.IsEven(val)
}

func IsOdd[T math.Integer](val T) bool {
	return math.IsOdd(val)
}

func Max[T cmp.Ordered](a, b T, rest ...T) T {
	return math.Max(a, b, rest...)
}

func Min[T cmp.Ordered](a, b T, rest ...T) T {
	return math.Min(a, b, rest...)
}

func Clamp[T cmp.Ordered](val, low, high T) T {
	return math.Clamp(val, low, high)
}

func Abs[T math.SignedNumber](val T) T {
	return math.Abs(val)
}

// Numeric functions
func Sum[T numeric.Number](slice []T) T {
	return numeric.Sum(slice)
}

func Product[T numeric.Number](slice []T) T {
	return numeric.Product(slice)
}

func Iota[T numeric.Integer](start T, size int) []T {
	return numeric.Iota(start, size)
}

func MaxElement[T cmp.Ordered](slice []T) (T, bool) {
	return numeric.MaxElement(slice)
}

func MinElement[T cmp.Ordered](slice []T) (T, bool) {
	return numeric.MinElement(slice)
}

func PrefixSum[T numeric.Number](slice []T) []T {
	return numeric.PrefixSum(slice)
}