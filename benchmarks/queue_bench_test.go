package benchmarks

import (
	"fmt"
	"testing"

	"github.com/mohitjoshi-hey/stlgo/container/queue"
)

func BenchmarkQueue_Enqueue(b *testing.B) {
	for _, n := range benchmarkSizes {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				q := queue.New[int]()
				for j := 0; j < n; j++ {
					q.Enqueue(j)
				}
			}
		})
	}
}

func BenchmarkQueue_Dequeue(b *testing.B) {
	for _, n := range benchmarkSizes {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				q := queue.New[int]()
				for j := 0; j < n; j++ {
					q.Enqueue(j)
				}
				for !q.IsEmpty() {
					q.Dequeue()
				}
			}
		})
	}
}

func BenchmarkQueue_InterleavedHotPath(b *testing.B) {
	q := queue.New[int]()
	b.ReportAllocs()
	i := 0
	for b.Loop() {
		q.Enqueue(i)
		if i%2 == 0 {
			q.Dequeue()
		}
		i++
	}
}