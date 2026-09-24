package sort

import (
	"slices"
	"testing"
)

func FuzzNthElement(f *testing.F) {
	f.Add([]byte{1, 5, 3, 2, 4}, 2)
	f.Add([]byte{7, 7, 7, 7}, 1)
	f.Add([]byte{9, 4, 1, 7, 2}, 0)

	f.Fuzz(func(t *testing.T, data []byte, n int) {
		if len(data) == 0 || n < 0 || n >= len(data) {
			return;
		}

		slice := make([]int, len(data))
		for i, b := range data {
			slice[i] = int(b);
		}

		want := slices.Clone(slice);
		slices.Sort(want);
		wantVal := want[n];

		got := NthElement(slice, n);

		if got != wantVal {
			t.Fatalf("NthElement(%v, %d) = %d, want %d", data, n, got, wantVal)
		}
		if slice[n] != wantVal {
			t.Fatalf("slice[%d] = %d after NthElement, want %d", n, slice[n], wantVal)
		}

		for i := 0; i < n; i++ {
			if slice[i] > slice[n] {
				t.Fatalf("left invariant broken at i=%d: slice[i]=%d > slice[n]=%d", i, slice[i], slice[n])
			}
		}
		for i := n + 1; i < len(slice); i++ {
			if slice[i] < slice[n] {
				t.Fatalf("right invariant broken at i=%d: slice[i]=%d < slice[n]=%d", i, slice[i], slice[n])
			}
		}
	})
}