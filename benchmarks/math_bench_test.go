package benchmarks

import (
	"testing"

	"github.com/mohitjoshi-hey/stlgo/algo"
)

func BenchmarkMath_GCD(b *testing.B) {
	b.ReportAllocs()
	var res int
	for b.Loop() {
		res = algo.GCD(1234567890, 987654321)
	}
	sinkInt = res
}

func BenchmarkMath_GCD_Variadic(b *testing.B) {
	b.ReportAllocs()
	var res int
	for b.Loop() {
		res = algo.GCD(120, 240, 360, 480, 600)
	}
	sinkInt = res
}

func BenchmarkMath_LCM(b *testing.B) {
	b.ReportAllocs()
	var res int
	for b.Loop() {
		res = algo.LCM(123456, 789012)
	}
	sinkInt = res
}

func BenchmarkMath_IsPrime(b *testing.B) {
	primeTarget := 1_000_000_007
	b.ReportAllocs()
	var res bool
	for b.Loop() {
		res = algo.IsPrime(primeTarget)
	}
	sinkBool = res
}

func BenchmarkMath_Max_Variadic(b *testing.B) {
	b.ReportAllocs()
	var res int
	for b.Loop() {
		res = algo.Max(12, 45, 78, 23, 99, 54, 102, 3)
	}
	sinkInt = res
}