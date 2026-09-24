package search

import (
	"slices"
	"testing"
)

func FuzzSearchBounds(f *testing.F) {
	f.Add([]byte{1, 2, 2, 5, 7}, byte(2))
	f.Add([]byte{}, byte(0))
	f.Add([]byte{9, 9, 9}, byte(9))

	f.Fuzz(func(t *testing.T, data []byte, target byte) {
		s := make([]int, len(data))
		for i, b := range data {
			s[i] = int(b);
		}
		slices.Sort(s);
		val := int(target)

		lb := LowerBound(s, val);
		ub := UpperBound(s, val);

		wantLB := 0
		for wantLB < len(s) && s[wantLB] < val {
			wantLB++;
		}
		wantUB := 0
		for wantUB < len(s) && s[wantUB] <= val {
			wantUB++;
		}

		if lb != wantLB {
			t.Fatalf("LowerBound(%v, %d) = %d, want %d", s, val, lb, wantLB)
		}
		if ub != wantUB {
			t.Fatalf("UpperBound(%v, %d) = %d, want %d", s, val, ub, wantUB)
		}
		if lb > ub {
			t.Fatalf("invariant broken: LowerBound=%d > UpperBound=%d for val=%d in %v", lb, ub, val, s)
		}

		idx, found := BinarySearch(s, val)
		wantFound := lb < len(s) && s[lb] == val
		if found != wantFound {
			t.Fatalf("BinarySearch(%v, %d) found=%v, want %v", s, val, found, wantFound)
		}
		if found && idx != lb {
			t.Fatalf("BinarySearch(%v, %d) = %d, want %d (LowerBound)", s, val, idx, lb)
		}
	})
}
