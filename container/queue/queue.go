// this package queue provides a generic FIFO queue.
// I made it as a thin wrapper over the deque, just to reduce some hardwork.
package queue

import "github.com/mohitjoshi-hey/stlgo/container/deque"

// Queue is a generic first-in-first-out queue backed by container/deque.
type Queue[T any] struct {
	d *deque.Deque[T]
}

// New creates a Queue pre-populated with items, front to back.
func New[T any](items ...T) *Queue[T] {
	return &Queue[T]{d: deque.New(items...)}
}

// Enqueue appends value to the back of the queue.
func (q *Queue[T]) Enqueue(value T) { 
	q.d.PushBack(value)
}

// Dequeue removes and returns the element at the front of the queue. ok is false if the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) { 
	return q.d.PopFront()
}

// GetFront returns the element at the front of the queue without removing it. ok is false if the queue is empty.
func (q *Queue[T]) GetFront() (T, bool) { 
	return q.d.GetFront()
}

// GetRear returns the element at the back of the queue without removing it. ok is false if the queue is empty.
func (q *Queue[T]) GetRear() (T, bool) { 
	return q.d.GetRear()
}

// IsEmpty reports whether the queue holds no elements.
func (q *Queue[T]) IsEmpty() bool{ 
	return q.d.IsEmpty()
}

// Len returns the number of elements currently stored.
func (q *Queue[T]) Len() int { 
	return q.d.Len()
}