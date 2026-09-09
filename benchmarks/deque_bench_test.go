package benchmarks

import (
	"fmt"
	"testing"

	"github.com/mohitjoshi-hey/stlgo/container/deque"
)

func BenchmarkDeque_PushBack(b *testing.B) {
	for _, n := range benchmarkSizes {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				dq := deque.New[int]()
				for j := 0; j < n; j++ {
					dq.PushBack(j)
				}
			}
		})
	}
}

func BenchmarkDeque_PushFront(b *testing.B) {
	for _, n := range benchmarkSizes {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				dq := deque.New[int]()
				for j := 0; j < n; j++ {
					dq.PushFront(j)
				}
			}
		})
	}
}

func BenchmarkDeque_CircularHotPath(b *testing.B) {
	dq := deque.New[int]()
	// Pre-fill to establish buffer
	for i := 0; i < 8; i++ {
		dq.PushBack(i)
	}

	b.ReportAllocs()
	for b.Loop() {
		dq.PushBack(1)
		dq.PushFront(2)
		dq.PopFront()
		dq.PopBack()
	}
}