// Package deque provides a generic double-ended queue backed by a circular ring buffer. Both ends support O(1) amortized push/pop, with automatic doubling and unwrap-on-grow once the buffer fills.
package deque

// Deque is a generic double-ended queue. head and tail are internal ring indices and must not be set directly from outside the package.
type Deque[T any] struct {
	val []T
	head int
	tail int
	len int
}

// New creates a Deque pre-populated with items (front to back). If fewer than 4 items are given, the backing array is still allocated with a minimum capacity of 4 to avoid an immediate reallocation on first push.
func New[T any](items ...T) *Deque[T] {
        // While starting empty, allocate a constant size(like 4) to avoid immediate resizing
	capacity := len(items);
	if capacity < 4 {
		capacity = 4;
	}

	d := &Deque[T]{
		val: make([]T, capacity),
		head: 0,
		tail: len(items) % capacity,
		len: len(items),
	}

	copy(d.val, items)
	return d;
}

// GetFront returns the element at the front of the deque without removing it. ok is false if the deque is empty.
func (d *Deque[T]) GetFront() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	return d.val[d.head], true
}

// GetRear returns the element at the back of the deque without removing it. ok is false if the deque is empty.
func (d *Deque[T]) GetRear() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	capacity := len(d.val)
	return d.val[(d.tail-1+capacity)%capacity], true
}

// IsEmpty reports whether the deque holds no elements.
func (d *Deque[T]) IsEmpty() bool {
	return d.len == 0;
}

// Len returns the number of elements currently stored.
func (d *Deque[T]) Len() int {
	return d.len;
}

func (d *Deque[T]) grow() {
	newCap := len(d.val)*2;
	newVal := make([]T, newCap)

	n := copy(newVal, d.val[d.head:])
	copy(newVal[n:], d.val[:d.head])

	d.val = newVal;
	d.head = 0;
	d.tail = d.len;
}

// PushBack appends value to the back of the deque in O(1) amortized time.
func (d *Deque[T]) PushBack(value T) {
	if d.len == len(d.val) {
		d.grow();
	}

	d.val[d.tail] = value
	d.tail = (d.tail + 1) % len(d.val)
	d.len++;
}

// PushFront prepends value to the front of the deque in O(1) amortized time.
func (d *Deque[T]) PushFront(value T) {
	if d.len == len(d.val) {
		d.grow();
	}

	d.head = (d.head - 1 + len(d.val)) % len(d.val)
	d.val[d.head] = value
	d.len++;
}

// PopBack removes and returns the element at the back of the deque. ok is false if the deque is empty.
func (d *Deque[T]) PopBack() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}

	capacity := len(d.val)
	d.tail = (d.tail - 1 + capacity) % capacity

	ele := d.val[d.tail]

	var zero T
	d.val[d.tail] = zero
	d.len--

	return ele, true
}

// PopFront removes and returns the element at the front of the deque. ok is false if the deque is empty.
func (d *Deque[T]) PopFront() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}

	ele := d.val[d.head]

	var zero T
	d.val[d.head] = zero

	d.head = (d.head + 1) % len(d.val)
	d.len--

	return ele, true
}