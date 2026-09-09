package benchmarks

import (
	"testing"

	"github.com/mohitjoshi-hey/stlgo/algo"
)

func BenchmarkPermutation_NextPermutation(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = algo.NextPermutation(arr)
	}
}

func BenchmarkPermutation_PrevPermutation(b *testing.B) {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = algo.PrevPermutation(arr)
	}
}