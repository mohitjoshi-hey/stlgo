package algo

import "cmp"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Sum[T Number](slice []T) T {
	sum := T(0)
	for _, val := range slice {
		sum += val
	}

	return sum
}

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