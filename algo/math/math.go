// Package math provides common competitive-programming numeric utilities- GCD, LCM, testing prime, Max/Min/Clamp, and simple integer predicates.
package math

import (
	"cmp"
)

// Integer covers Go's standard signed and unsigned integer types.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// SignedNumber covers signed integer and floating-point types.
type SignedNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func gcdTwo[T Integer](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

// GCD returns the greatest common divisor of a, b, and any additional values in rest. The result is always non-negative.
func GCD[T Integer](a, b T, rest ...T) T {
	res := gcdTwo(a, b)
	if res == 1 {
		return 1
	}
	for _, v := range rest {
		res = gcdTwo(res, v)
		if res == 1 {
			break
		}
	}
	return res
}

func lcmTwo[T Integer](a, b T) T {
	if a == 0 || b == 0 {
		return 0
	}
	g := gcdTwo(a, b)
	res := (a / g) * b
	if res < 0 {
		return -res
	}
	return res
}

// LCM returns the least common multiple of a, b, and any additional values in rest. Returns 0 if any argument is 0.
func LCM[T Integer](a, b T, rest ...T) T {
	res := lcmTwo(a, b)
	if res == 0 {
		return 0
	}
	for _, v := range rest {
		res = lcmTwo(res, v)
		if res == 0 {
			break
		}
	}
	return res
}

// IsPrime checks primality in O(sqrt(n)) time without integer overflow.
func IsPrime[T Integer](n T) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := T(5); i <= n/i; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func max2[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Max returns the largest of a, b, and any additional values in rest.
func Max[T cmp.Ordered](a, b T, rest ...T) T {
	m := max2(a, b)
	for _, ele := range rest {
		m = max2(m, ele)
	}
	return m
}

func min2[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Min returns the smallest of a, b, and any additional values in rest.
func Min[T cmp.Ordered](a, b T, rest ...T) T {
	m := min2(a, b)
	for _, ele := range rest {
		m = min2(m, ele)
	}
	return m
}

// IsEven tells whether val is evenly divisible by 2.
func IsEven[T Integer](val T) bool {
	return val%2 == 0
}

// IsOdd tells whether val is not evenly divisible by 2.
func IsOdd[T Integer](val T) bool {
	return !IsEven(val)
}

// Abs returns the absolute value of val.
func Abs[T SignedNumber](val T) T {
	if val < 0 {
		return -val;
	}
	return val;
}

func Clamp[T cmp.Ordered](val, low, high T) T {
	if high < low {
		panic("math: Clamp called with low > high");
	}
	if val < low {
		return low;
	}
	if val > high {
		return high;
	}

	return val
}