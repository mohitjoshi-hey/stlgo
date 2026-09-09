package benchmarks

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/mohitjoshi-hey/stlgo/algo"
)

func BenchmarkSort_NthElement_Vs_SlicesSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 50_000} {
		k := n / 2
		source := make([]int, n)
		for i := 0; i < n; i++ {
			source[i] = rand.IntN(1_000_000)
		}

		b.Run(fmt.Sprintf("NthElement_N=%d", n), func(b *testing.B) {
			buf := make([]int, n)
			b.ReportAllocs()
			for b.Loop() {
				copy(buf, source)
				_, _ = algo.NthElement(buf, k)
			}
		})

		b.Run(fmt.Sprintf("SlicesSort_N=%d", n), func(b *testing.B) {
			buf := make([]int, n)
			b.ReportAllocs()
			for b.Loop() {
				copy(buf, source)
				slices.Sort(buf)
			}
		})
	}
}