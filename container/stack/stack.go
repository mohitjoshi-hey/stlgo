package stack

type Stack[T any] struct {
	value []T
}

func New[T any](items ...T) *Stack[T] {
	q := &Stack[T]{
		value:  make([]T, 0, len(items)),
	}

	q.value = append(q.value, items...)
	return q
}

func (s *Stack[T]) Push(val T) {
	s.value = append(s.value, val)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.value) == 0
}

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

func (s *Stack[T]) Len() int {
	return len(s.value);
}

func (s *Stack[T]) Top() (T, bool) {
	if s.Len() == 0 {
		var zero T
		return zero, false
	}

	index := s.Len() - 1;
	ele := s.value[index];
	return ele, true;
}
