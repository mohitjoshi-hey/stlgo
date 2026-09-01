package queue

type Queue[T any] struct {
	val  []T
	head int
}

func New[T any](items ...T) *Queue[T] {
	q := &Queue[T]{
		val:  make([]T, 0, len(items)), 
		head: 0,
	}
	q.val = append(q.val, items...)
	return q
}

func (q *Queue[T]) Enqueue(value T) {
	q.val = append(q.val, value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	value := q.val[q.head]

	var zero T
	q.val[q.head] = zero
	q.head++

	if q.head == len(q.val) {
		q.val = q.val[:0]
		q.head = 0
	} else if q.head > 100 && q.head >= len(q.val)/2 {
		copy(q.val, q.val[q.head:])
		q.val = q.val[:len(q.val)-q.head]
		q.head = 0
	}

	return value, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.val)-q.head == 0
}

func (q *Queue[T]) Len() int {
	return len(q.val) - q.head
}

func (q *Queue[T]) GetFront() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.val[q.head], true
}

func (q *Queue[T]) GetRear() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.val[len(q.val)-1], true
}