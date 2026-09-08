package permutations

import (
	"slices"
	"testing"
)

func TestPermutations(t *testing.T) {
	t.Run("Testing NextPermutation Function", func(t *testing.T) {
		s := []int{2, 3, 5, 4, 1}

		got, ok := NextPermutation(s);
		want := []int{2, 4, 1, 3, 5}
		if !ok || !slices.Equal(got, want) {
			t.Errorf("NextPermutations gives Error, got (%v, %t), want (%v, true)", got, ok, want)
		}
	})

	t.Run("NextPermutation Last Permutation", func(t *testing.T) {
		s := []int{5, 4, 3, 2, 1}

		got, ok := NextPermutation(s);
		want := []int{1, 2, 3, 4, 5}
		if ok || !slices.Equal(got, want) {
			t.Errorf("NextPermutations gives Error, got (%v, %t), want (%v, false)", got, ok, want)
		}
	})

	t.Run("Testing PrevPermutation Function", func(t *testing.T) {
		s := []int{2, 3, 5, 4, 1}

		got, ok := PrevPermutation(s);
		want := []int{2, 3, 5, 1, 4}
		if !ok || !slices.Equal(got, want) {
			t.Errorf("PrevPermutations gives Error, got (%v, %t), want (%v, true)", got, ok, want)
		}
	})

	t.Run("PrevPermutation First Permutation", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}

		got, ok := PrevPermutation(s);
		want := []int{5, 4, 3, 2, 1}
		if ok || !slices.Equal(got, want) {
			t.Errorf("PrevPermutations gives Error, got (%v, %t), want (%v, false)", got, ok, want)
		}
	})

	t.Run("Checking PrevPermutations & NextPermutations simultaneously", func(t *testing.T) {
		s := []int{2, 3, 5, 4, 1}
		want :=slices.Clone(s);
		
		PrevPermutation(s);
		got, ok := NextPermutation(s);
		if !ok || !slices.Equal(got, want) {
			t.Errorf("PrevPermutation followed by NextPermutation did not return original state")
		}
	})
}