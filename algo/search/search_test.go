package search

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("Test LowerBound", func(t *testing.T) {
		s := []int{1, 2, 5, 7, 23, 23, 45}
		
		if ind := LowerBound(s, 23); ind != 4 {
			t.Errorf("LowerBound function is wrong, expected 4, got %d", ind)
		}
		
		if ind := LowerBound(s, 0); ind != 0 {
			t.Errorf("Expected 0 for target smaller than all elements, got %d", ind)
		}
	})

	t.Run("Test UpperBound", func(t *testing.T) {
		s := []int{1, 2, 5, 7, 23, 23, 45}
		
		if ind := UpperBound(s, 23); ind != 6 {
			t.Errorf("UpperBound function is wrong, expected 6, got %d", ind)
		}
		if ind := UpperBound(s, 99); ind != len(s) {
			t.Errorf("Expected %d for target larger than all elements, got %d", len(s), ind)
		}
	})

	t.Run("Test BinarySearch", func(t *testing.T) {
		s := []int{1, 2, 2, 5, 7, 23, 23, 45}
		
		ind, found := BinarySearch(s, 5)
		if !found || ind != 3 {
			t.Errorf("BinarySearch expected (3, true), got (%d, %t)", ind, found)
		}

		ind, found = BinarySearch(s, 99)
		if found || ind != -1 {
			t.Errorf("BinarySearch expected (-1, false), got (%d, %t)", ind, found)
		}
	})

	t.Run("Test Generic Strings", func(t *testing.T) {
		s := []string{"apple", "banana", "cherry", "date"}
		
		ind, found := BinarySearch(s, "cherry")
		if !found || ind != 2 {
			t.Errorf("Generic BinarySearch expected (2, true) for string, got (%d, %t)", ind, found)
		}
	})

	t.Run("Test Empty Slice", func(t *testing.T) {
		var s []int
		
		if ind := LowerBound(s, 5); ind != 0 {
			t.Errorf("LowerBound on empty slice should return 0, got %d", ind)
		}
		if ind := UpperBound(s, 5); ind != 0 {
			t.Errorf("UpperBound on empty slice should return 0, got %d", ind)
		}
		if _, found := BinarySearch(s, 5); found {
			t.Errorf("BinarySearch on empty slice should return false, got true")
		}
	})

	t.Run("Test the Find function", func(t *testing.T) {
		slice := []int{23, 12, 3, 56, 32}

		ind, found := Find(slice, 12)
		if ind != 1 || !found {
			t.Errorf("Find failed on valid element, expected (1, true), got (%d, %t)", ind, found)
		}

		ind, found = Find(slice, 67)
		if ind != -1 || found {
			t.Errorf("Find failed on missing element, expected (-1, false), got (%d, %t)", ind, found)
		}
	})

	t.Run("Test the FindIf function", func(t *testing.T) {
		nums := []int{1, 3, 7, 12, 19, 24}
		idx, found := FindIf(nums, func(v int) bool {
			return v%2 == 0
		})
		if !found || idx != 3 {
			t.Errorf("FindIf expected (3, true) for first even number, got (%d, %t)", idx, found)
		}

		idx, found = FindIf(nums, func(v int) bool {
			return v > 100
		})
		if found || idx != -1 {
			t.Errorf("FindIf expected (-1, false) for unmatched condition, got (%d, %t)", idx, found)
		}

		type Item struct {
			Name  string
			Price int
		}
		items := []Item{
			{Name: "Pen", Price: 10},
			{Name: "Book", Price: 50},
			{Name: "Laptop", Price: 800},
		}

		idx, found = FindIf(items, func(item Item) bool {
			return item.Price >= 50
		})
		if !found || idx != 1 {
			t.Errorf("FindIf expected (1, true) for struct predicate, got (%d, %t)", idx, found)
		}
	})
}