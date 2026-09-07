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
		if !IsEven(22) || IsEven(25) {
			t.Errorf("IsEven function has a fault, expected true, got false");
		}

		if IsOdd(22) || !IsOdd(25) {
			t.Errorf("IsOdd function has a fault, expected false, got true");
		}
	})

	t.Run("Testing Abs Function", func(t *testing.T) {
		if Abs(3) != 3 {
			t.Errorf("Abs int failed, expected 3, got %d", Abs(3))
		}
		if Abs(-3) != 3 {
			t.Errorf("Abs int failed, expected 3, got %d", Abs(3))
		}
	})

	t.Run("Testing Clamp Function", func(t *testing.T) {
		if got := Clamp(3, 5, 9); got != 5 {
			t.Errorf("Clamp failed, got %d, expected 5", got)
		}
		if got := Clamp(10, 5, 9); got != 9 {
			t.Errorf("Clamp failed, got %d, expected 9", got)
		}
		if got := Clamp(7, 5, 9); got != 7 {
			t.Errorf("Clamp failed, got %d, expected 7", got)
		}
	})

	t.Run("Testing Clamp Panic on high < low", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil { // If the code doesnot panics then it will return nil, otherwise !nil. 
				t.Errorf("Clamp did not panic on low > high")
			}
		}()
		Clamp(10, 11, 9) // It's gonna fail.
	})
}