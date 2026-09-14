package sort

import (
	"slices"
	"testing"
)

func TestSortAndSelect(t *testing.T) {
	t.Run("Testing the NthElement function", func(t *testing.T) {
		original := []int{42, 17, 93, 8, 55, 23, 67, 31, 88, 12, 76, 5, 99, 34, 61}

		sorted := make([]int, len(original))
		copy(sorted, original)
		slices.Sort(sorted)

		for k := 0; k < len(original); k++ {
			trial := make([]int, len(original))
			copy(trial, original)

			val := NthElement(trial, k)

			//Checking exact value
			if val != sorted[k] || trial[k] != sorted[k] {
				t.Fatalf("Index %d: expected %d, got val=%d, trial[k]=%d", k, sorted[k], val, trial[k])
			}

			//Check left invariant (<= pivot)
			for i := 0; i < k; i++ {
				if trial[i] > trial[k] {
					t.Fatalf("Invariant broken at k=%d: left element trial[%d]=%d > pivot=%d", k, i, trial[i], trial[k])
				}
			}

			// Check right invariant (>= pivot)
			for i := k + 1; i < len(trial); i++ {
				if trial[i] < trial[k] {
					t.Fatalf("Invariant broken at k=%d: right element trial[%d]=%d < pivot=%d", k, i, trial[i], trial[k])
				}
			}
		}
	})

	t.Run("Checking if, will the elements of slice gets aligned if an element is inserted at a particular index of the slice", func(t *testing.T) {
		s := []int{9, 4, 1, 7, 2}
		k := 2
		val := NthElement(s, k)

		if val != 4 || s[k] != 4 {
			t.Errorf("Expected 4 at index %d, got %d", k, val)
		}
	})

	t.Run("Testing a SortBy with a Custom Structs", func(t *testing.T) {
		type Player struct {
			Name  string
			Score int
		}

		players := []Player{
			{Name: "Charlie", Score: 70},
			{Name: "Alice", Score: 95},
			{Name: "Bob", Score: 85},
		}

		// Sort by Score descending
		SortBy(players, func(a, b Player) bool {
			return a.Score > b.Score
		})

		expected := []string{"Alice", "Bob", "Charlie"}
		for i, p := range players {
			if p.Name != expected[i] {
				t.Errorf("SortBy failed at index %d: expected %s, got %s", i, expected[i], p.Name)
			}
		}
	})

	t.Run("Testing ouut of  bounds edge cases", func(t *testing.T) {
		s := []int{1, 2, 3}

		expectPanic := func(fn func()) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic, got none")
				}
			}()
			fn()
		}

		expectPanic(func() {
			 NthElement(s, -1)
		})
		expectPanic(func() {
			 NthElement(s, 5)
		})

		var empty []int
		expectPanic(func() {
			 NthElement(empty, 0)
		})
	})

	t.Run("Testing the IsSorted function", func(t *testing.T) {
		s1 := []int{9, 4, 1, 7, 2}
		s2 := []int{1, 2, 3, 4, 5}

		if IsSorted(s1) {
			t.Errorf("Flaw in IsSorted, expected false, got true")
		}
		if !IsSorted(s2) {
			t.Errorf("Flaw in IsSorted, expected true, got false")
		}
	})

	t.Run("All Identical Elements (A Killer Test for Lomuto)", func(t *testing.T) {
		s := []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7}
		val := NthElement(s, 4)
		if val != 7 || s[4] != 7 {
			t.Errorf("Expected 7 at index 4, got %d", val)
		}
	})
}