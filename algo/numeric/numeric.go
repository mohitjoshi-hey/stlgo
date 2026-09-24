// Package numeric provides generic arithmetic and array-generation helpers: sums, products, ranges, extrema, and prefix sums.
package numeric

import "cmp"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

// Integer covers Go's standard signed and unsigned integer types.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Sum returns the sum of all elements in slice. Returns 0 for an empty slice.
func Sum[T Number](slice []T) T {
	sum := T(0)
	for _, val := range slice {
		sum += val
	}

	return sum
}

// Product returns the product of all elements in slice. Returns 0 for an empty slice (there is nothing to multiply).
func Product[T Number](slice []T) T {
	product := T(1)
	if len(slice) == 0 {
		return T(0)
	}
	for _, val := range slice {
		product *= val
	}

	return product
}

// Iota returns a slice of size consecutive values starting at start. Returns nil if size <= 0.
func Iota[T Integer](start T, size int) []T {
	if size <= 0 {
		return nil
	}

	arr := make([]T, size)
	input := start
	for i := 0; i < size; i++ {
		arr[i] = input
		input++
	}
	return arr
}

// MaxElement returns the largest element in slice. ok is false if slice is empty.
func MaxElement[T cmp.Ordered](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}

	max := slice[0];
	for _, val := range slice[1:] {
		if val > max {
			max = val;
		}
	}
	return max, true
}

// MinElement returns the smallest element in slice. ok is false if slice is empty.
func MinElement[T cmp.Ordered](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}

	min := slice[0];
	for _, val := range slice[1:] {
		if val < min {
			min = val;
		}
	}
	return min, true
}

// PrefixSum returns a new slice where each element is the cumulative sum of slice up to and including that index. Returns nil for an empty slice.
func PrefixSum[T Number](slice []T) []T {
	if len(slice) == 0 {
		return nil;
	}

	preSum := make([]T, len(slice))
	preSum[0] = slice[0];
	for i, val := range slice {
		if i > 0 {
			preSum[i] = preSum[i - 1] + val;
		}
	}
	return preSum;
}