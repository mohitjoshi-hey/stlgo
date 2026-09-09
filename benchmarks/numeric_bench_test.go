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

func BenchmarkNumeric_MinMaxElement(b *testing.B) {
	for _, n := range benchmarkSizes {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = i * 3
		}

		b.Run(fmt.Sprintf("MaxElement_N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			var res int
			for b.Loop() {
				res, _ = algo.MaxElement(arr)
			}
			sinkInt = res
		})
	}
}