package benchmarks

import (
	"fmt"
	"testing"

	"github.com/mohitjoshi-hey/stlgo/algo"
)

func BenchmarkNumeric_PrefixSum(b *testing.B) {
	for _, n := range benchmarkSizes {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = i + 1
		}

		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			var res []int
			for b.Loop() {
				res = algo.PrefixSum(arr)
			}
			sinkInts = res
		})
	}
}

func BenchmarkNumeric_Sum(b *testing.B) {
	for _, n := range benchmarkSizes {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = i + 1
		}

		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			var res int
			for b.Loop() {
				res = algo.Sum(arr)
			}
			sinkInt = res
		})
	}
}

func BenchmarkNumeric_Product(b *testing.B) {
	for _, n := range benchmarkSizes {
		// Keep values small (1 or 2) so the product benchmark measures loop/
		// multiply overhead rather than overflow behaviour on large N.
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = 1 + (i % 2)
		}

		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			var res int
			for b.Loop() {
				res = algo.Product(arr)
			}
			sinkInt = res
		})
	}
}

func BenchmarkNumeric_Iota(b *testing.B) {
	for _, n := range benchmarkSizes {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			var res []int
			for b.Loop() {
				res = algo.Iota(1, n)
			}
			sinkInts = res
		})
	}
}

// BenchmarkNumeric_MaxElement and BenchmarkNumeric_MinElement are kept as separate benchmark functions (rather than one "MinMaxElement" function that only exercised Max) so both code paths actually get measured and show up as distinct rows in `go test -bench`.
func BenchmarkNumeric_MaxElement(b *testing.B) {
	for _, n := range benchmarkSizes {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = i * 3
		}

		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			var res int
			for b.Loop() {
				res, _ = algo.MaxElement(arr)
			}
			sinkInt = res
		})
	}
}

func BenchmarkNumeric_MinElement(b *testing.B) {
	for _, n := range benchmarkSizes {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = i * 3
		}

		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			var res int
			for b.Loop() {
				res, _ = algo.MinElement(arr)
			}
			sinkInt = res
		})
	}
}