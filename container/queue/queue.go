// this package queue provides a generic FIFO queue.
// I made it as a thin wrapper over the deque, just to reduce some hardwork.
package queue

import "github.com/mohitjoshi-hey/stlgo/container/deque"

type Queue[T any] struct {
	d *deque.Deque[T]
}

func New[T any](items ...T) *Queue[T] {
	return &Queue[T]{d: deque.New(items...)}
}

func (q *Queue[T]) Enqueue(value T) { 
	q.d.PushBack(value)
}
func (q *Queue[T]) Dequeue() (T, bool) { 
	return q.d.PopFront()
}
func (q *Queue[T]) GetFront() (T, bool) { 
	return q.d.GetFront()
}
func (q *Queue[T]) GetRear() (T, bool) { 
	return q.d.GetRear()
}
func (q *Queue[T]) IsEmpty() bool{ 
	return q.d.IsEmpty()
}
func (q *Queue[T]) Len() int { 
	return q.d.Len()
}