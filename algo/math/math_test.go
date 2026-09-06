package math

import "testing"

func TestMathFunctions(t *testing.T) {
	t.Run("Check LCM Functionn", func(t *testing.T) {
		got := LCM(5, 4, 11, 8);
		want := 440;

		if got != want {
			t.Errorf("LCM function shows error, got %d, expected %d", got, want);
		}
	})

	t.Run("Check GCD Functionn", func(t *testing.T) {
		got := GCD(5, 20, 100, 15);
		want := 5;

		if got != want {
			t.Errorf("GCD function shows error, got %d, expected %d", got, want);
		}
	})

	t.Run("Testing IsPrime Functionn", func(t *testing.T) {
		got := IsPrime(23);
		want := true;
		if got != want {
			t.Errorf("IsPrime function has a fault, expected true, got false");
		}

		got = IsPrime(99);
		want = false;
		if got != want {
			t.Errorf("IsPrime function has a fault, expected false, got true");
		}
	})

	t.Run("Testing Max & Min Functionn", func(t *testing.T) {
		got := Max(23, 4, 67, 12);
		want := 67;
		if got != want {
			t.Errorf("Max function has a fault, expected %d, got %d", want, got);
		}

		got = Min(23, 4, 67, 12);
		want = 4;
		if got != want {
			t.Errorf("Min function has a fault, expected %d, got %d", want, got);
		}
	})

	t.Run("Testing IsEven & IsOdd Functionn", func(t *testing.T) {
		got := IsEven(22);
		want := true;
		if got != want {
			t.Errorf("IsEven function has a fault, expected true, got false");
		}

		got = IsEven(23);
		want = false;
		if got != want {
			t.Errorf("IsEven function has a fault, expected false, got true");
		}
	})
}