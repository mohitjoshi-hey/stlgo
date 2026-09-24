// Package stack provides a generic, slice-backed LIFO stack.
package stack

// Stack is a generic last-in-first-out stack.
type Stack[T any] struct {
	value []T
}

// New creates a Stack pre-populated with items, bottom to top — the last item given ends up on top.
func New[T any](items ...T) *Stack[T] {
	q := &Stack[T]{
		value:  make([]T, 0, len(items)),
	}

	q.value = append(q.value, items...)
	return q
}

// Push adds val to the top of the stack.
func (s *Stack[T]) Push(val T) {
	s.value = append(s.value, val)
}

// IsEmpty reports whether the stack holds no elements.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.value) == 0
}

// Pop removes and returns the element at the top of the stack. ok is false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.value) - 1
	ele := s.value[index]
	
	var zero T
	s.value[index] = zero 
	
	s.value = s.value[:index]

	return ele, true
}

// Len returns the number of elements currently stored.
func (s *Stack[T]) Len() int {
	return len(s.value);
}

// Top returns the element at the top of the stack without removing it. ok is false if the stack is empty.
func (s *Stack[T]) Top() (T, bool) {
	if s.Len() == 0 {
		var zero T
		return zero, false
	}

	index := s.Len() - 1;
	ele := s.value[index];
	return ele, true;
}