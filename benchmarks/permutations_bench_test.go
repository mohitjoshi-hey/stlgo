package benchmarks

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/mohitjoshi-hey/stlgo/algo"
)

var permutationSizes = []int{5, 10, 12}

// newShuffledBase returns a randomly ordered slice of size n. Starting from a random order (rather than the fully ascending/descending array) avoids always hitting the cheapest possible pivot position, which is what a sorted or reverse-sorted starting slice gives you.
func newShuffledBase(n int) []int {
	base := make([]int, n)
	for i := range base {
		base[i] = i
	}
	rand.Shuffle(n, func(i, j int) { base[i], base[j] = base[j], base[i] })
	return base
}

func BenchmarkPermutation_NextPermutation(b *testing.B) {
	for _, n := range permutationSizes {
		base := newShuffledBase(n)
		arr := make([]int, n)

		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				// Reset to the same starting arrangement each iteration so every call measures a consistent, reproducible input instead of drifting through the whole permutation space (which previously made the benchmark's cost depend on how many iterations b.Loop() happened to run).
				//
				copy(arr, base)
				_, _ = algo.NextPermutation(arr)
			}
		})
	}
}

func BenchmarkPermutation_PrevPermutation(b *testing.B) {
	for _, n := range permutationSizes {
		base := newShuffledBase(n)
		arr := make([]int, n)

		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				copy(arr, base)
				_, _ = algo.PrevPermutation(arr)
			}
		})
	}
}