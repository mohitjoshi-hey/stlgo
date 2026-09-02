package deque

import (
	"slices"
	"testing"
)

func TestDequeOperations(t *testing.T) {
	t.Run("Initialization and Edge Cases", func(t *testing.T) {
		d := New(10, 20)

		if d.Len() != 2 {
			t.Errorf("Expected len 2, got %d", d.Len())
		}
		if d.IsEmpty() {
			t.Error("Expected deque to not be empty")
		}

		d.PopFront()
		d.PopFront()

		if _, ok := d.PopFront(); ok {
			t.Error("PopFront on empty deque should return false")
		}
		if _, ok := d.PopBack(); ok {
			t.Error("PopBack on empty deque should return false")
		}
	})

	t.Run("PushBack and PopFront (Queue Behavior)", func(t *testing.T) {
		d := New[int]()
		d.PushBack(1)
		d.PushBack(2)
		d.PushBack(3)

		v, ok := d.PopFront()
		if !ok || v != 1 {
			t.Errorf("Expected 1, got %d", v)
		}
		if d.Len() != 2 {
			t.Errorf("Expected len 2, got %d", d.Len())
		}
	})

	t.Run("PushFront and PopBack (Stack Behavior)", func(t *testing.T) {
		d := New[int]()
		d.PushFront(1)
		d.PushFront(2)
		d.PushFront(3)

		v, ok := d.PopBack()
		if !ok || v != 1 {
			t.Errorf("Expected 1, got %d", v)
		}
	})

	t.Run("Circular Wrapping Logic", func(t *testing.T) {
		d := New[int]() // Base capacity is 4

		// Fill capacity: [1, 2, 3, _]
		d.PushBack(1)
		d.PushBack(2)
		d.PushBack(3)

		// Pop one from front: [_, 2, 3, _], Head moves to index 1
		d.PopFront()

		// Push two more: [5, 2, 3, 4]
		// Tail wraps around to index 0 using modulo!
		d.PushBack(4)
		d.PushBack(5)

		// Verify the internal array state explicitly
		if d.Head != 1 || d.Tail != 1 {
			t.Errorf("Head and Tail should both be 1 due to wrapping, got Head: %d, Tail: %d", d.Head, d.Tail)
		}
		if !slices.Equal(d.val, []int{5, 2, 3, 4}) {
			t.Errorf("Circular wrap failed, got internal array: %v", d.val)
		}
	})

	t.Run("Dynamic Resizing (Grow)", func(t *testing.T) {
		d := New[int]() // Base capacity is 4

		// Create a wrapped state: [5, 2, 3, 4]
		d.PushBack(1)
		d.PushBack(2)
		d.PushBack(3)
		d.PopFront() // Removes 1
		d.PushBack(4)
		d.PushBack(5) 

		// Push one more. Capacity is 4, Len is 4. This must trigger grow()
		d.PushBack(6)

		// The circular array should have been "unwrapped" into a new capacity-8 array
		if len(d.val) != 8 {
			t.Errorf("Expected capacity 8, got %d", len(d.val))
		}
		if d.Head != 0 {
			t.Errorf("Expected Head to reset to 0 after grow, got %d", d.Head)
		}
		
		// Expected sequence from front to back is 2, 3, 4, 5, 6
		expectedFront := []int{2, 3, 4, 5, 6}
		for i, exp := range expectedFront {
			v, _ := d.PopFront()
			if v != exp {
				t.Errorf("After grow, expected element %d at pop %d, got %d", exp, i, v)
			}
		}
	})
}