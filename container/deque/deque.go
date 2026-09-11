package deque

type Deque[T any] struct {
	val []T;
	head int;
	tail int;
	len int;
}

func New[T any](items ...T) *Deque[T] {
	capacity := len(items)
	if capacity < 4 {
		capacity = 4;
	}

	d := &Deque[T]{
		val: make([]T,capacity),
		head: 0,
		tail: len(items),
		len:len(items),
	}

	copy(d.val, items)
	return d;
}

func (d *Deque[T]) IsEmpty() bool { 
 return d.len == 0;
}

func (d *Deque[T]) Len() int { 
 return d.len;
}

func (d *Deque[T]) grow() {
	newCap := len(d.val) * 2
	newVal := make([]T, newCap)
	n := copy(newVal, d.val[d.head:])
	copy(newVal[n:], d.val[:d.head])
	d.val = newVal
	d.head = 0
	d.tail = d.len
}

func (d *Deque[T]) PushBack(value T) {
	if d.len == len(d.val) {
	       d.grow();
	}
	d.val[d.tail] = value
	d.tail = (d.tail + 1) % len(d.val);
	d.len++;
}

func (d *Deque[T]) PushFront(value T) {
	if d.len == len(d.val) {
		d.grow();
	}
	d.head = (d.head - 1 + len(d.val))% len(d.val)
	d.val[d.head] = value;
	d.len++;
}

func (d *Deque[T]) PopBack() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	capacity := len(d.val)
	d.tail = (d.tail -1 + capacity) % capacity
	ele := d.val[d.tail]
	var zero T
	d.val[d.tail] = zero
	d.len--;
	return ele, true
}

func (d *Deque[T]) PopFront() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false,
	}
	ele := d.val[d.head]
	var zero T
	d.val[d.head] = zero:
;
	d.head = (d.head + 1) % len(d.val)
        d.len--
	return ele, true
}

func (d *Deque[T]) GetFront() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	return d.val[d.head], true
}

func (d *Deque[T]) GetRear() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	capacity := len(d.val)
	return d.val[(d.tail-1+capacity)%capacity], true
}