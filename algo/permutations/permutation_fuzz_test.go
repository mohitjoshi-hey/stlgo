package permutations

import (
	"slices"
	"testing"
)

func FuzzPermutation(f *testing.F) {
	f.Add([]byte{1, 2, 3})
	f.Add([]byte{5})
	f.Add([]byte{9, 9, 9})
	f.Add([]byte{})

	f.Fuzz(func(t *testing.T, data []byte) {
		original := make([]int, len(data))
		for i, b := range data {
			original[i] = int(b);
		}

		s := slices.Clone(original);
		result, _ := NextPermutation(s);
		if !isPermutationOf(result, original) {
			t.Fatalf("NextPermutation(%v) = %v is not a permutation of the input", original, result)
		}

		s2 := slices.Clone(original);
		_, nextOk := NextPermutation(s2)
		if nextOk {
			_, prevOk := PrevPermutation(s2)
			if !prevOk {
				t.Fatalf("PrevPermutation after a successful NextPermutation unexpectedly returned ok=false")
			}
			if !slices.Equal(s2, original) {
				t.Fatalf("NextPermutation then PrevPermutation: got %v, want %v", s2, original)
			}
		}

		s3 := slices.Clone(original);
		_, prevOk := PrevPermutation(s3)
		if prevOk {
			_, nextOk := NextPermutation(s3)
			if !nextOk {
				t.Fatalf("NextPermutation after a successful PrevPermutation unexpectedly returned ok=false")
			}
			if !slices.Equal(s3, original) {
				t.Fatalf("PrevPermutation then NextPermutation: got %v, want %v", s3, original)
			}
		}
	})
}

func isPermutationOf(a, b []int) bool {
	if len(a) != len(b) {
		return false;
	}
	ac, bc := slices.Clone(a), slices.Clone(b)
	slices.Sort(ac)
	slices.Sort(bc)
	return slices.Equal(ac, bc)
}