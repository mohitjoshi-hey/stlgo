package benchmarks

import (
	"fmt"
	"testing"

	"github.com/mohitjoshi-hey/stlgo/container/stack"
)

func BenchmarkStack_Push(b *testing.B) {
	for _, n := range benchmarkSizes {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				st := stack.New[int]()
				for j := 0; j < n; j++ {
					st.Push(j)
				}
			}
		})
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	for _, n := range benchmarkSizes {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				st := stack.New[int]()
				for j := 0; j < n; j++ {
					st.Push(j)
				}
				for !st.IsEmpty() {
					st.Pop()
				}
			}
		})
	}
}

func BenchmarkStack_HotPath_PushPop(b *testing.B) {
	st := stack.New[int]()
	b.ReportAllocs()
	for b.Loop() {
		st.Push(42)
		st.Pop()
	}
}