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
		d := New[int]() // base capacity is 4

		// Fill capacity: [1, 2, 3, _]
		d.PushBack(1)
		d.PushBack(2)
		d.PushBack(3)

		
		d.PopFront()

		d.PushBack(4)
		d.PushBack(5)

		if d.head != 1 || d.tail != 1 {
			t.Errorf("Head and Tail should both be 1 due to wrapping, got Head: %d, Tail: %d", d.head, d.tail)
		}
		if !slices.Equal(d.val, []int{5, 2, 3, 4}) {
			t.Errorf("Circular wrap failed, got internal array: %v", d.val)
		}
	})

	t.Run("Testing the grow funjtion", func(t *testing.T) {
		d := New[int]()

		d.PushBack(1)
		d.PushBack(2)
		d.PushBack(3)
		d.PopFront()
		d.PushBack(4)
		d.PushBack(5) 

		// We will puush one more. Capacity is 4, Len is 4. This must trigger grow()
		d.PushBack(6)

		// The circular array should have been "unwrapped" into a new capacity-8 array
		if len(d.val) != 8 {
			t.Errorf("Expected capacity 8, got %d", len(d.val))
		}
		if d.Head != 0 {
			t.Errorf("Expected Head to reset to 0 after grow, got %d", d.Head)
		}
		
		expectedFront := []int{2, 3, 4, 5, 6}
		for i, exp := range expectedFront {
			v, _ := d.PopFront()
			if v != exp {
				t.Errorf("After grow, expected element %d at pop %d, got %d", exp, i, v)
			}
		}
	})
}