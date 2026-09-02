package deque

type Deque[T any] struct {
	val []T
	Head int
	Tail int
	len int
}

func New[T any](items ...T) *Deque[T] {
	// If starting empty, allocate a small base capacity (e.g., 4) to avoid immediate resizing
	cap := len(items)
	if cap < 4 {
		cap = 4
	}

	d := &Deque[T]{
		val: make([]T, cap),
		Head: 0,
		Tail: len(items),
		len: len(items),
	}

	copy(d.val, items)
	return d
}

func (d *Deque[T]) IsEmpty() bool {
	return d.len == 0
}

func (d *Deque[T]) Len() int {
	return d.len
}

// grow doubles the capacity and "unwraps" the circular buffer
func (d *Deque[T]) grow() {
	newCap := len(d.val) * 2
	newVal := make([]T, newCap)

	// Because the buffer is 100% full, Head and Tail are at the exact same index.
	// We unwrap by copying from Head to the end, then from the start up to Head.
	n := copy(newVal, d.val[d.Head:])
	copy(newVal[n:], d.val[:d.Head]) // We can just use Head here since Head == Tail

	d.val = newVal
	d.Head = 0
	d.Tail = d.len // Tail points to the next empty slot
}

func (d *Deque[T]) PushBack(value T) {
	if d.len == len(d.val) {
		d.grow()
	}

	d.val[d.Tail] = value
	d.Tail = (d.Tail + 1) % len(d.val)
	d.len++
}

func (d *Deque[T]) PushFront(value T) {
	if d.len == len(d.val) {
		d.grow()
	}

	d.Head = (d.Head - 1 + len(d.val)) % len(d.val)
	d.val[d.Head] = value
	d.len++
}

func (d *Deque[T]) PopBack() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}

	// Move tail backward to find the last inserted element
	cap := len(d.val)
	d.Tail = (d.Tail - 1 + cap) % cap

	ele := d.val[d.Tail]

	// Erase reference to prevent memory leaks
	var zero T
	d.val[d.Tail] = zero
	d.len--

	return ele, true
}

func (d *Deque[T]) PopFront() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}

	ele := d.val[d.Head]

	// Erase reference to prevent memory leaks
	var zero T
	d.val[d.Head] = zero

	// Move Head forward
	d.Head = (d.Head + 1) % len(d.val)
	d.len--

	return ele, true
}