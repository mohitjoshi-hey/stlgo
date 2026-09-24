// Package algo is a facade over stlgo's algorithm subpackages (algo/search, algo/sort, algo/math, algo/numeric, algo/permutations),
// re-exporting each function under one import path. Each function below only forwards to its underlying implementation — see the linked subpackage for full documentation.
package algo

import (
	"cmp"

	math "github.com/mohitjoshi-hey/stlgo/algo/math"
	numeric "github.com/mohitjoshi-hey/stlgo/algo/numeric"
	permutation "github.com/mohitjoshi-hey/stlgo/algo/permutations"
	search "github.com/mohitjoshi-hey/stlgo/algo/search"
	sort "github.com/mohitjoshi-hey/stlgo/algo/sort"
)

// Search functions

// LowerBound; see search.LowerBound for details.
func LowerBound[T cmp.Ordered](slice []T, val T) int {
	return search.LowerBound(slice, val)
}

// UpperBound; see search.UpperBound for details.
func UpperBound[T cmp.Ordered](slice []T, val T) int {
	return search.UpperBound(slice, val)
}

// BinarySearch; see search.BinarySearch for details.
func BinarySearch[T cmp.Ordered](slice []T, val T) (int, bool) {
	return search.BinarySearch(slice, val)
}

// Find; see search.Find for details.
func Find[T comparable](slice []T, val T) (int, bool) {
	return search.Find(slice, val)
}

// FindIf; see search.FindIf for details.
func FindIf[T any](slice []T, pred func(val T) bool) (int, bool) {
	return search.FindIf(slice, pred)
}

// Sort functions

// Sort; see sort.Sort for details.
func Sort[T cmp.Ordered](slice []T) {
	sort.Sort(slice)
}

// DescSort; see sort.DescSort for details.
func DescSort[T cmp.Ordered](slice []T) {
	sort.DescSort(slice)
}

// SortBy; see sort.SortBy for details.
func SortBy[T any](slice []T, less func(a, b T) bool) {
	sort.SortBy(slice, less)
}

// IsSorted; see sort.IsSorted for details.
func IsSorted[T cmp.Ordered](slice []T) bool {
	return sort.IsSorted(slice)
}

// NthElement panics if n < 0 or n >= len(slice); see sort.NthElement for details.
func NthElement[T cmp.Ordered](slice []T, n int) T {
	return sort.NthElement(slice, n)
}

// Math functions

// GCD; see math.GCD for details.
func GCD[T math.Integer](a, b T, rest ...T) T {
	return math.GCD(a, b, rest...)
}

// LCM; see math.LCM for details.
func LCM[T math.Integer](a, b T, rest ...T) T {
	return math.LCM(a, b, rest...)
}

// IsPrime; see math.IsPrime for details.
func IsPrime[T math.Integer](n T) bool {
	return math.IsPrime(n)
}

// IsEven; see math.IsEven for details.
func IsEven[T math.Integer](val T) bool {
	return math.IsEven(val)
}

// IsOdd; see math.IsOdd for details.
func IsOdd[T math.Integer](val T) bool {
	return math.IsOdd(val)
}

// Max; see math.Max for details.
func Max[T cmp.Ordered](a, b T, rest ...T) T {
	return math.Max(a, b, rest...)
}

// Min; see math.Min for details.
func Min[T cmp.Ordered](a, b T, rest ...T) T {
	return math.Min(a, b, rest...)
}

// Clamp panics if low > high; see math.Clamp for details.
func Clamp[T cmp.Ordered](val, low, high T) T {
	return math.Clamp(val, low, high)
}

// Abs; see math.Abs for details.
func Abs[T math.SignedNumber](val T) T {
	return math.Abs(val)
}

// Numeric functions

// Sum; see numeric.Sum for details.
func Sum[T numeric.Number](slice []T) T {
	return numeric.Sum(slice)
}

// Product; see numeric.Product for details.
func Product[T numeric.Number](slice []T) T {
	return numeric.Product(slice)
}

// Iota; see numeric.Iota for details.
func Iota[T numeric.Integer](start T, size int) []T {
	return numeric.Iota(start, size)
}

// MaxElement; see numeric.MaxElement for details.
func MaxElement[T cmp.Ordered](slice []T) (T, bool) {
	return numeric.MaxElement(slice)
}

// MinElement; see numeric.MinElement for details.
func MinElement[T cmp.Ordered](slice []T) (T, bool) {
	return numeric.MinElement(slice)
}

// PrefixSum; see numeric.PrefixSum for details.
func PrefixSum[T numeric.Number](slice []T) []T {
	return numeric.PrefixSum(slice)
}

// Permutation functions

// NextPermutation; see permutation.NextPermutation for details.
func NextPermutation[T cmp.Ordered](slice []T) ([]T, bool) {
	return permutation.NextPermutation(slice)
}

// PrevPermutation; see permutation.PrevPermutation for details.
func PrevPermutation[T cmp.Ordered](slice []T) ([]T, bool) {
	return permutation.PrevPermutation(slice)
}