package benchmarks

import (
	"fmt"
	"testing"

	"github.com/mohitjoshi-hey/stlgo/algo"
)

func BenchmarkSearch_LowerBound(b *testing.B) {
	for _, n := range benchmarkSizes {
		data := make([]int, n)
		for i := 0; i < n; i++ {
			data[i] = i * 2
		}
		target := n

		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			var idx int
			for b.Loop() {
				idx = algo.LowerBound(data, target)
			}
			sinkInt = idx
		})
	}
}

func BenchmarkSearch_UpperBound(b *testing.B) {
	for _, n := range benchmarkSizes {
		data := make([]int, n)
		for i := 0; i < n; i++ {
			data[i] = i * 2
		}
		target := n

		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			var idx int
			for b.Loop() {
				idx = algo.UpperBound(data, target)
			}
			sinkInt = idx
		})
	}
}

func BenchmarkSearch_BinarySearch(b *testing.B) {
	for _, n := range benchmarkSizes {
		data := make([]int, n)
		for i := 0; i < n; i++ {
			data[i] = i * 2
		}
		target := n

		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			var idx int
			var ok bool
			for b.Loop() {
				idx, ok = algo.BinarySearch(data, target)
			}
			sinkInt = idx
			sinkBool = ok
		})
	}
}

func BenchmarkSearch_Find_Linear(b *testing.B) {
	for _, n := range benchmarkSizes {
		data := make([]int, n)
		for i := 0; i < n; i++ {
			data[i] = i
		}
		target := n - 1

		b.Run(fmt.Sprintf("WorstCase_N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			var idx int
			for b.Loop() {
				idx, _ = algo.Find(data, target)
			}
			sinkInt = idx
		})
	}
}

func BenchmarkSearch_FindIf(b *testing.B) {
	for _, n := range benchmarkSizes {
		data := make([]int, n)
		for i := 0; i < n; i++ {
			data[i] = i
		}
		target := n - 1

		b.Run(fmt.Sprintf("Predicate_N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			var idx int
			for b.Loop() {
				idx, _ = algo.FindIf(data, func(v int) bool {
					return v == target
				})
			}
			sinkInt = idx
		})
	}
}