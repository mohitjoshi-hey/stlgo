package stack

type Stack[T any] struct {
	value []T
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
		return zero, false;
	}

	index := len(s.value) - 1;
	ele := s.value[index]
	s.value = s.value[:index]

	return ele, true;
}

func (s *Stack[T]) Len() int {
	return len(s.value);
}

