package deque

import "testing"

func FuzzDeque(f *testing.F) {
	f.Add([]byte{}, []byte{1, 2, 3})
	f.Add([]byte{0, 1, 2, 3}, []byte{})
	f.Add([]byte{0, 0, 1, 3, 2, 0, 1, 1, 2, 3}, []byte{9})

	f.Fuzz(func(t *testing.T, ops []byte, seedBytes []byte) {
		if len(ops) > 500 || len(seedBytes) > 100 {
			t.Skip();
		}

		seed := make([]int, len(seedBytes))
		for i, b := range seedBytes {
			seed[i] = int(b);
		}

		d := New(seed...)
		ref := append([]int(nil), seed...)

		pushVal := 0
		for _, op := range ops {
			switch op % 4 {
			case 0:
				pushVal++
				d.PushFront(pushVal)
				ref = append([]int{pushVal}, ref...)
			case 1:
				pushVal++
				d.PushBack(pushVal);
				ref = append(ref, pushVal)
			case 2:
				gotVal, gotOk := d.PopFront()
				wantOk := len(ref) > 0
				if gotOk != wantOk {
					t.Fatalf("PopFront ok mismatch: got %v, want %v", gotOk, wantOk)
				}
				if wantOk {
					wantVal := ref[0];
					ref = ref[1:];
					if gotVal != wantVal {
						t.Fatalf("PopFront value mismatch: got %d, want %d", gotVal, wantVal)
					}
				}
			case 3: // PopBack
				gotVal, gotOk := d.PopBack();
				wantOk := len(ref) > 0
				if gotOk != wantOk {
					t.Fatalf("PopBack ok mismatch: got %v, want %v", gotOk, wantOk)
				}
				if wantOk {
					wantVal := ref[len(ref)-1]
					ref = ref[:len(ref)-1]
					if gotVal != wantVal {
						t.Fatalf("PopBack value mismatch: got %d, want %d", gotVal, wantVal)
					}
				}
			}

			if d.Len() != len(ref) {
				t.Fatalf("Len mismatch: got %d, want %d (ref=%v)", d.Len(), len(ref), ref)
			}
			if d.IsEmpty() != (len(ref) == 0) {
				t.Fatalf("IsEmpty mismatch: got %v, want %v", d.IsEmpty(), len(ref) == 0)
			}

			frontVal, frontOk := d.GetFront()
			wantFrontOk := len(ref) > 0
			if frontOk != wantFrontOk {
				t.Fatalf("GetFront ok mismatch: got %v, want %v", frontOk, wantFrontOk)
			}
			if wantFrontOk && frontVal != ref[0] {
				t.Fatalf("GetFront value mismatch: got %d, want %d", frontVal, ref[0])
			}

			rearVal, rearOk := d.GetRear()
			wantRearOk := len(ref) > 0
			if rearOk != wantRearOk {
				t.Fatalf("GetRear ok mismatch: got %v, want %v", rearOk, wantRearOk)
			}
			if wantRearOk && rearVal != ref[len(ref)-1] {
				t.Fatalf("GetRear value mismatch: got %d, want %d", rearVal, ref[len(ref)-1])
			}
		}
	})
}