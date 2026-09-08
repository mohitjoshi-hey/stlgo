package algo

import (
	"slices"
	"testing"
)

func TestNumeric(t *testing.T) {
	t.Run("Test the Sum function", func(t *testing.T) {

		nums := []int{10, 20, 30, 40, 50}
		if got := Sum(nums[1:4]); got != 90 {
			t.Errorf("Sum(nums[1:4]) = %d; want 90", got)
		}

		if got := Sum([]int{-5, 10, -3, 2}); got != 4 {
			t.Errorf("Sum() = %d; want 4", got)
		}

		if got := Sum([]int{}); got != 0 {
			t.Errorf("Sum([]) = %d; want 0", got)
		}
	})

	t.Run("Test the Product function", func(t *testing.T) {
		nums := []int{-2, 3, -23, 9, -4}

		if got := Product([]int{-2, 3, -4}); got != 24 {
			t.Errorf("Product() = %d; want 24", got)
		}
		if got := Product(nums[1:3]); got != -69 {
			t.Errorf("Product() = %d; want -69", got)
		}
		if got := Product([]int{}); got != 0 {
			t.Errorf("Product([]) = %d; want 0", got)
		}
	})

	t.Run("Test the Product function", func(t *testing.T) {
		got := Iota(5, 4)
		want := []int{5, 6, 7, 8}
		
		if !slices.Equal(got, want) {
			t.Errorf("Iota(5, 4) = %v; want %v", got, want)
		}

		gotNeg := Iota(-2, 5)
		wantNeg := []int{-2, -1, 0, 1, 2}
		if !slices.Equal(gotNeg, wantNeg) {
			t.Errorf("Iota(-2, 5) = %v; want %v", gotNeg, wantNeg)
		}

		if got := Iota(1, 0); got != nil {
			t.Errorf("Iota(1, 0) = %v; want nil", got)
		}
		if got := Iota(1, -5); got != nil {
			t.Errorf("Iota(1, -5) = %v; want nil", got)
		}
	})

	t.Run("MaxElement", func(t *testing.T) {

		mixed := []int{10, -5, 20, 15}
		val, ok := MaxElement(mixed)
		if !ok || val != 20 {
			t.Errorf("MaxElement(mixed) = (%d, %t); want (20, true)", val, ok)
		}

		words := []string{"apple", "zebra", "banana"}
		sVal, sOk := MaxElement(words)
		if !sOk || sVal != "zebra" {
			t.Errorf("MaxElement(words) = (%s, %t); want (\"zebra\", true)", sVal, sOk)
		}

		if _, ok := MaxElement([]int{}); ok {
			t.Errorf("MaxElement([]) ok = true; want false")
		}
	})

	t.Run("MinElement", func(t *testing.T) {
		neg := []int{-30, -10, -50, 5, -20}
		val, ok := MinElement(neg)
		if !ok || val != -50 {
			t.Errorf("MinElement(neg) = (%d, %t); want (-50, true)", val, ok)
		}

		if _, ok := MinElement([]int{}); ok {
			t.Errorf("MinElement([]) ok = true; want false")
		}
	})

	t.Run("PrefixSum", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := PrefixSum(nums)
		want := []int{1, 3, 6, 10, 15}
		if !slices.Equal(got, want) {
			t.Errorf("PrefixSum() = %v; want %v", got, want)
		}

		mixed := []int{2, -3, 5, -1, 4}
		got= PrefixSum(mixed)
		want= []int{2, -1, 4, 3, 7}
		
		if !slices.Equal(got, want) {
			t.Errorf("PrefixSum(mixed) = %v; want %v", got, want)
		}
		if gotEmpty := PrefixSum([]int{}); gotEmpty != nil {
			t.Errorf("PrefixSum([]) = %v; want nil", gotEmpty)
		}
	})
}