package stack

import (
	"slices"
	"testing"
)

func TestStackOperations(t *testing.T) {
	t.Run("Formation of Stack", func(t *testing.T) {
		st1 := New[int]()
		st2 := New(1, 2, 3)

		if !st1.IsEmpty() {
			t.Errorf("st1 expected to be empty")
		}
		if st2.IsEmpty() {
			t.Errorf("st2 expected not to be empty")
		}
		if st2.Len() != 3 {
			t.Errorf("st2 length is expected to be %d, got %d", 3, st2.Len())
		}
		if !slices.Equal(st2.value, []int{1, 2, 3}) {
			t.Errorf("Expected {1, 2, 3}, got %v", st2)
		}
	})
	t.Run("Push & Pop Operations", func(t *testing.T) {
		st1 := New[int]()

		st1.Push(10)
		st1.Push(20)
		st1.Push(30)

		if st1.IsEmpty() {
			t.Errorf("Push doesn't work, expected {10, 20, 30}, got %v", st1)
		}
		if !slices.Equal(st1.value, []int{10, 20, 30}) {
			t.Errorf("Push doesn't work, expected {10, 20, 30}, got %v", st1)
		}

		val, ok := st1.Pop()

		if val != 30 {
			t.Errorf("Pop doesn't work, expected 30, got %d", val)
		}
		if !ok {
			t.Errorf("Expected ok = true")
		}
		if !slices.Equal(st1.value, []int{10, 20}) {
			t.Errorf("Pop doesn't work, expected {10, 20}, got %v", st1)
		}

		_, _ = st1.Pop()
		_, _ = st1.Pop()
		_, ok = st1.Pop()

		if ok {
			t.Errorf("Expected ok = false")
		}
	})

	t.Run("Testing Other Operations", func(t *testing.T) {
		st1 := New[int]()
		st2 := New(1, 2, 3)

		if !st1.IsEmpty() {
			t.Errorf("IsEmpty() not working, expected true")
		}
		if st2.IsEmpty() {
			t.Errorf("IsEmpty() not working, expected false")
		}
		if st2.Len() != 3 {
			t.Errorf("Len() not working, expected %d, got %d", 3, st2.Len())
		}

		_, ok1 := st1.Top()
		if ok1 { 
			t.Errorf("Top() on empty stack should return ok = false")
		}

		val2, ok2 := st2.Top()
		if !ok2 { 
			t.Errorf("Top() on populated stack should return ok = true")
		}
		if val2 != 3 {
			t.Errorf("Top() not working, expected 3, got %d", val2)
		}
	})
}
